package data

import (
	"fmt"
	"github.com/spf13/viper"
	"go-music-uniapp/server/db"
	"log"
	"testing"
)

func TestDb_to_csv(t *testing.T) {
	initViper()
	pg := db.InitPgDB()
	defer func() {
		err := pg.Close()
		if err != nil {
			log.Println("An error occurred while database was closing the connection : ", err)
		}
	}()
	DbToCsv()
}

func initViper() {
	fmt.Printf("Loading configuration logics...\n")
	//workDir, _ := os.Getwd() //获取工作目录
	//viper := viper.New()		//不要New	即可全局使用
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to get the configuration.")
	}
}
