package dataGetter

import (
	"crypto/md5"
	"fmt"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/model"
	"go-music-uniapp/server/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

func CreateUsers(num int, startTelephone, initPwd string) {
	for i := 0; i < num; i++ {
		username := util.RandomString(10)
		phoneInt, _ := strconv.Atoi(startTelephone)
		telephone := fmt.Sprintf("%d", phoneInt+i)
		md5PwdByte := md5.Sum([]byte(initPwd))
		md5PwdStr := fmt.Sprintf("%x", md5PwdByte)                                          //模拟前端md5加密后的密码
		hasedPwd, err := bcrypt.GenerateFromPassword([]byte(md5PwdStr), bcrypt.DefaultCost) //后端再对md5加密后的密码进行hash加密
		if err != nil {
			log.Println("加密错误")
			return
		}
		newUser := model.User{Username: username, Telephone: telephone, Password: string(hasedPwd)}
		//fmt.Printf("readData 读到数据 = %v \n", newUser)
		db.PGEngine.Create(&newUser)
		// 每次发送完时等待
		time.Sleep(time.Second)
	}
}

func CreateFavorList() {

}
