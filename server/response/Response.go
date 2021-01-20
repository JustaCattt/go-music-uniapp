package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//	成功
func Success(ctx *gin.Context, data gin.H, msg string)  {
	Response(ctx, http.StatusOK, 200, data, msg)
}

//	业务失败
func Fail(ctx *gin.Context, data gin.H, msg string)  {
	Response(ctx, http.StatusOK, 400, data, msg)
}

//	用户验证失败
func ValidFail(ctx *gin.Context, msg string)  {
	Response(ctx, http.StatusUnprocessableEntity, 422, nil, msg)
}

//	服务器错误
func ServerError(ctx *gin.Context, msg string)  {
	Response(ctx, http.StatusInternalServerError, 500, nil, msg)
}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string)  {
	//	data为nil则不返回data
	if data == nil {
		ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg})
		return
	}
	//	msg为""则不返回msg
	if msg == "" {
		ctx.JSON(httpStatus, gin.H{"code": code, "data": data})
		return
	}
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}