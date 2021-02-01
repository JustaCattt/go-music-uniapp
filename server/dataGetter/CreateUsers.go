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
		md5PwdStr := fmt.Sprintf("%x", md5PwdByte) //模拟前端md5加密后的密码
		//fmt.Printf("md5加密后的密码：%v\n", md5PwdStr)
		hasedPwd, err := bcrypt.GenerateFromPassword([]byte(md5PwdStr), bcrypt.DefaultCost) //后端再对md5加密后的密码进行hash加密
		//fmt.Printf("hash加密后的密码：%v\n", string(hasedPwd))
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

type Result struct {
	Id   int
	Name string
}

func CreateFavorList(n int) {
	var users []model.User
	db.PGEngine.Find(&users)
	var results []Result
	db.PGEngine.Table("songs").Select("id, name").Find(&results)
	for i := len(users) - 1; i >= 0; i-- {
		fmt.Printf("userId:%d\n", users[i].ID)
		for j := 0; j < n; j++ {
			key := util.RandomNum(len(results))
			fmt.Printf("随机索引为%d\t", key)
			fmt.Printf("songId:%v\t", results[key].Id)
			fmt.Printf("歌名:《%v》\t", results[key].Name)
			count := util.RandomNum(100)
			fmt.Printf("播放次数为%d\t", count)
			newFavorList := model.FavorList{
				UserId:    users[i].ID,
				SongName:  results[key].Name,
				SongId:    results[key].Id,
				PlayCount: count,
			}
			fmt.Printf("favorlist:%v\n", newFavorList)
			db.PGEngine.Create(&newFavorList)
			time.Sleep(time.Second)
		}
	}
	fmt.Printf("\n---已为%d位用户，生成数量为%d的收藏列表---\n", len(users), n)
}
