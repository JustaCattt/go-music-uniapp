package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-music-uniapp/server/data"
	"go-music-uniapp/server/dataGetter"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/router"
	"log"
	"os"
)

func main() {
	initViper()
	pg := db.InitPgDB()
	defer func() {
		err := pg.Close()
		if err != nil {
			log.Println("An error occurred while database was closing the connection : ", err)
		}
	}()
	if viper.GetBool("first_start") {
		fmt.Println("开始获取数据...")
		dataGetter.CreateUsers(20, "19968086600", "123456") //创建20个用户
		dataGetter.InitData()                               //获取歌曲数据
		dataGetter.CreateFavorList(20)                      //为每个用户随机生成收藏歌单
		data.DbToCsv()
		fmt.Println("获取数据完毕")
	} else {
		router.InitRouter()
	}
}

func initViper() {
	fmt.Printf("Loading configuration logics...\n")
	workDir, _ := os.Getwd() //获取工作目录
	//viper := viper.New()		//不要New	即可全局使用
	viper.SetConfigName("config")
	viper.AddConfigPath(workDir)
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to get the configuration.")
	}
}
