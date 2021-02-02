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
	"time"
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

		dataGetter.InitData()                               //获取歌曲数据，用时不确定，快的话一分钟足以，因为启用了goroutine
		time.Sleep(3 * time.Minute)							//这里阻塞是为了保证后面的表有数据依赖，前三分钟获取歌单，后7分钟制造假数据并生成csv

		dataGetter.CreateUsers(20, "19968086600", "123456") //创建20个用户，20s
		dataGetter.CreateFavorList(20)                      //为每个用户随机生成收藏歌单,20*20s=400s，加起来总共7分钟

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
