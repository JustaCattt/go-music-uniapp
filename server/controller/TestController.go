package controller

import (
	"github.com/gin-gonic/gin"
	"go-music-uniapp/server/response"
)

func Test(ctx *gin.Context) {
	response.Success(ctx, nil, "接口正常")
}
