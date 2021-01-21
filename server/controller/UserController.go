package controller

import (
	"github.com/gin-gonic/gin"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/model"
	"go-music-uniapp/server/response"
	"go-music-uniapp/server/util"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Register(ctx *gin.Context) {
	//	获取参数
	username := ctx.PostForm("username")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//	数据验证
	if len(telephone) != 11 {
		response.ValidFail(ctx, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.ValidFail(ctx, "密码不能小于6位")
		return
	}

	//	如果用户名没传，给一个10位的随机字符串
	if len(username) == 0 {
		username = util.RandomString(10)
	}

	//	判断手机号是否存在
	var user model.User
	db.PGEngine.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		response.ValidFail(ctx, "手机号已被注册")
		return
	}

	//	创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.ServerError(ctx, "加密错误")
		return
	}
	newUser := model.User{
		Username:  username,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.PGEngine.Create(&newUser)

	//	返回结果
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	//	获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//	数据验证
	if len(telephone) != 11 {
		response.ValidFail(ctx, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.ValidFail(ctx, "密码不能小于6位")
		return
	}

	//	判断手机号是否存在
	var user model.User
	db.PGEngine.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.ValidFail(ctx, "用户不存在")
		return
	}

	//	判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.ValidFail(ctx, "密码错误")
		return
	}

	//	发放token
	token, err := util.ReleaseToken(user)
	if err != nil {
		response.ServerError(ctx, "系统异常")
		log.Printf("token generate error : %v\n", err)
		return
	}

	//	返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	//	返回结果
	response.Success(ctx, gin.H{"user": model.ToUserInfo(user.(model.User))}, "")
}
