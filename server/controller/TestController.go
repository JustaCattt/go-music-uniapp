package controller

import (
	"Music/server/response"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context)  {
	response.Success(ctx,nil,"接口正常")
}
