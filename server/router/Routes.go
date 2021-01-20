package router

import (
	"Music/server/controller"
	"Music/server/middleware"
	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) *gin.Engine {
	//接口测试路由
	r.GET("/test", controller.Test)

	//用户验证路由
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
		auth.GET("/info", middleware.AuthMiddleware(), controller.Info)
	}

	//返回路由集
	return r
}
