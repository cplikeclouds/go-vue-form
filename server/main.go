package main

import (
	"fmt"
	"server/core"
	"server/global"
	"server/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	//初始化日志

	global.Log = core.InitLogger()
	//连接数据库
	
	global.DB = core.InitGorm()
	//初始化路由
	r := routers.InitRouters()

	addr := global.Config.Server.AddHP()
	global.Log.Info(fmt.Sprintf("server listen at %s", addr))

	r.Run(addr)
}
