package middleware

import (
	"Music/server/db"
	"Music/server/model"
	"Music/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//	用户验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//	validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer "){
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims ,err := util.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足：无效token"})
			ctx.Abort()
			return
		}

		//	验证通过后获取claim中的userId
		userId := claims.UserId
		var user model.User
		db.PGEngine.First(&user, userId)

		//	用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户已注销"})
			ctx.Abort()
			return
		}

		//	用户存在	将用户信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}