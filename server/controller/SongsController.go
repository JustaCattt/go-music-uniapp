package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-music-uniapp/server/dataGetter"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/model"
	"go-music-uniapp/server/response"
	"strconv"
)

const RECOMMENDER_API = "http://127.0.0.1:8080"		//推荐服务
const NUMBER = 10		//推荐数量

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

//播放，对应地址为http://localhost:8090/song/play/:id
func Play(ctx *gin.Context)  {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//获取歌曲id
	id := ctx.Param("id")
	songId, _ := strconv.Atoi(id)

	//查询最近列表中是否存在
	var rp model.RecentlyPlayed
	db_result := db.PGEngine.Table("recently_playeds").Where("song_id = ?", songId).First(&rp)

	//若找不到
	if db_result.RecordNotFound() {
		var song model.Song
		db.PGEngine.Where("id = ?", songId).First(&song)
		newRp := model.RecentlyPlayed{
			UserId:    u.ID,
			SongName:  song.Name,
			SongId:    songId,
			PlayCount: 1,
		}
		db.PGEngine.Table("recently_playeds").Create(&newRp)
	} else {
		newpc := rp.PlayCount + 1
		db.PGEngine.Table("recently_playeds").Update("play_count", newpc)
	}

	var song model.Song
	db.PGEngine.Table("songs").Where("id = ?", songId).First(&song)
	var artist model.Artist
	db.PGEngine.Table("artists").Where("id = ?", song.ArtistId).First(&artist)
	playUrl := fmt.Sprintf("http://music.163.com/song/media/outer/url?id=%d", songId)
	result := Result{
		SongId:     songId,
		SongName:   song.Name,
		ArtistName: artist.Name,
		PlayUrl:    playUrl,
	}

	//返回结果
	response.Success(ctx, gin.H{"song":result}, "正在播放")
}

func Favor(ctx *gin.Context)  {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//获取歌曲id
	id := ctx.Param("id")
	songId, _ := strconv.Atoi(id)

	var fl model.FavorList
	db_result1 := db.PGEngine.Table("favor_lists").Where("song_id = ?", songId).First(&fl)

	//如果没有被收藏
	if db_result1.RecordNotFound() {
		//查询最近列表中是否存在
		var rp model.RecentlyPlayed
		db_result := db.PGEngine.Table("recently_playeds").Where("song_id = ?", songId).First(&rp)

		//若找不到
		if db_result.RecordNotFound() {
			var song model.Song
			db.PGEngine.Where("id = ?", songId).First(&song)
			newFavor := model.FavorList{
				UserId:    u.ID,
				SongName:  song.Name,
				SongId:    songId,
				PlayCount: 0,
			}
			db.PGEngine.Table("recently_playeds").Create(&newFavor)
		} else {
			newFavor := model.FavorList{
				UserId:    u.ID,
				SongName:  rp.SongName,
				SongId:    songId,
				PlayCount: rp.PlayCount,
			}
			db.PGEngine.Table("favor_lists").Create(&newFavor)
		}

		//返回结果
		response.Success(ctx, nil, "收藏成功")
	} else {	//如果收藏了
		db.PGEngine.Table("favor_lists").Where("song_id = ?", songId).Delete(&fl)

		//返回结果
		response.Success(ctx, nil, "取消收藏成功")
	}


}