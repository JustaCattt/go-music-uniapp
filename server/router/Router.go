package router

import (
	"Music/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())               //配置跨域中间件
	r = UseRoutes(r)                       //使用路由组
	port := viper.GetString("server.port") //获取配置信息
	if port != "" {
		panic(r.Run(":"+port))
	}
	panic(r.Run())
}
