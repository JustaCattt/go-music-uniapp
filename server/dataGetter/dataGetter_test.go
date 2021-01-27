package dataGetter

import (
	"fmt"
	"github.com/spf13/viper"
	"go-music-uniapp/server/db"
	"log"
	"testing"
)

func TestDataGetter(t *testing.T) {
	var artists = GetArtists()
	for _, artist := range artists {
		var songs = GetAllSongsByArtistId(artist.Id)
		fmt.Printf("\n%s全部的歌(共%d首)：\n", artist.Name, len(songs))
		for index, song := range songs {
			fmt.Printf("《%s》、", song.Name)
			if index%10 == 0 && index != 0 {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

func TestMap(t *testing.T) {
	m := MapArtistToSongs()
	fmt.Println(m["薛之谦"])
}

func TestCreateUsers(t *testing.T) {
	CreateUsers(20, "19968086600", "123456") //创建20个用户
}

func TestCreateFavorList(t *testing.T) {
	initViper()
	pg := db.InitPgDB()
	defer func() {
		err := pg.Close()
		if err != nil {
			log.Println("An error occurred while database was closing the connection : ", err)
		}
	}()
	CreateFavorList()
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
