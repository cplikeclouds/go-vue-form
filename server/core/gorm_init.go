package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/global"

	"log"
	"time"
)

func InitGorm() *gorm.DB {
	//检测配置HOST是否为空
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql,取消gorm连接")
		return nil
	}
	//gorm连接的字符串
	dsn := global.Config.Mysql.Dsn()

	//配置log模式
	var mysqlLogger logger.Interface
	if global.Config.Server.Env == "debug" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	//gorm连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s]Mysql链接失败", dsn)) //报错并且退出
	} else {
		fmt.Printf("[%s]Mysql链接成功\n", dsn)
	}

	//改到配置文件
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout

	return db
}
