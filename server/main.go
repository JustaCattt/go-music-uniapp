package main

import (
	"Music/server/db"
	"Music/server/router"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	initViper()
	pg := db.InitPgDB()
	defer pg.Close()
	router.InitRouter()
}

func initViper() {
	fmt.Printf("Loading configuration logics...\n")
	workDir, _ := os.Getwd()	//获取工作目录
	//viper := viper.New()		//不要New	即可全局使用
	viper.SetConfigName("config")
	viper.AddConfigPath(workDir)
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to get the configuration.")
	}
}