package routers

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouters() *gin.Engine {
	gin.SetMode(global.Config.Server.Env)
	r := gin.Default()

	privateGroup := r.Group("api/v1")
	//privateGroup.Use(middleware.JWTAuth())
	{
		//获取数据
		privateGroup.POST("/submit", v1.HandleFormSubmit)
		privateGroup.GET("/admin/datalist", v1.GetFormList)
		//privateGroup.GET("/admin/edit", v1.EditList)
		//privateGroup.GET("/admin/edit", v1.delete:id)
	}
	publicGroup := r.Group("api/v1")
	{
		//登录接口
		publicGroup.POST("login")
	}

	return r
}
