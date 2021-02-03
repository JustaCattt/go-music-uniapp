package db_to_csv

import (
	"fmt"
	"github.com/spf13/viper"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/util"
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
	DbToCsv2()
}

func DbToCsv2() {
	var results []Result
	db.PGEngine.Table("favor_lists").Find(&results)
	var data [][]string
	for index, _ := range results {
		fmt.Println(results[index])
		result := results[index]
		uIdStr := fmt.Sprintf("%d", result.UserId)
		sIdStr := fmt.Sprintf("%d", result.SongId)
		pcStr := fmt.Sprintf("%d", result.PlayCount)
		data = append(data, []string{uIdStr, sIdStr, pcStr})
	}
	util.CsvWriter("./data.csv", data)
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
