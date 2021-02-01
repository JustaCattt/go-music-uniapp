package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-music-uniapp/server/dataGetter"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/model"
	"go-music-uniapp/server/response"
)

const RECOMMENDER_API = "http://127.0.0.1:8080"
const NUMBER = 10

type Recommend struct {
	ItemId     string
	Popularity int
	Timestamp  string
	Score      int
}

type Result struct {
	SongId     int    `json:"id"`
	SongName   string `json:"name"`
	ArtistName string `json:"artist"`
	PlayUrl    string `json:"url"`
}

func Recommender(ctx *gin.Context) {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//拼接推荐服务的请求url，并解析到rs中
	url := fmt.Sprintf("%s/recommends/%d?number=%d", RECOMMENDER_API, u.ID, NUMBER)
	var rs []Recommend
	dataGetter.HttpGetParser(url, &rs)
	var results []Result

	//遍历查询
	for _, r := range rs {
		song := new(model.Song)
		artist := new(model.Artist)
		db.PGEngine.Table("songs").Where("id = ?", r.ItemId).First(&song)
		db.PGEngine.Table("artists").Where("id = ?", song.ArtistId).First(&artist)
		playUrl := fmt.Sprintf("http://music.163.com/song/media/outer/url?id=%d", song.Id)
		result := Result{
			SongId:     song.Id,
			SongName:   song.Name,
			ArtistName: artist.Name,
			PlayUrl:    playUrl,
		}
		results = append(results, result)
	}

	//打印推荐歌单
	for _, result := range results {
		fmt.Println(result)
	}

	//返回结果
	response.Success(ctx, gin.H{"songs": results}, "")
}
