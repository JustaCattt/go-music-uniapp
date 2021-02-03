package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/model"
	"go-music-uniapp/server/mygorse"
	"go-music-uniapp/server/response"
	"strconv"
)

type Result struct {
	SongId     int
	SongName   string
	ArtistName string
	PlayUrl    string
}

//推荐歌曲列表
func Recommender(ctx *gin.Context) {
	number := ctx.Param("number")
	num, _ := strconv.Atoi(number)

	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//获取推荐歌单
	rs := mygorse.GetRecommends(u.ID, num)

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
	msg := fmt.Sprintf("推荐歌曲%d首", num)
	response.Success(ctx, gin.H{"list": results}, msg)
}

//播放，对应地址为http://localhost:8090/song/play/:id
func Play(ctx *gin.Context) {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//获取歌曲id
	id := ctx.Param("id")
	songId, _ := strconv.Atoi(id)

	//查询最近列表中是否存在
	var rp model.RecentList
	db_result := db.PGEngine.Table("recent_lists").Where("song_id = ?", songId).First(&rp)

	//若找不到
	if db_result.RecordNotFound() {
		var song model.Song
		db.PGEngine.Where("id = ?", songId).First(&song)
		newRp := model.RecentList{
			UserId:    u.ID,
			SongName:  song.Name,
			SongId:    songId,
			PlayCount: 1,
		}
		db.PGEngine.Table("recent_lists").Create(&newRp)
	} else {
		newpc := rp.PlayCount + 1
		db.PGEngine.Table("recent_lists").Where(&model.RecentList{
			UserId: u.ID,
			SongId: songId,
		}).Update("play_count", newpc)
		db.PGEngine.Table("favor_lists").Where(&model.FavorList{
			UserId: u.ID,
			SongId: songId,
		}).Update("play_count", newpc)
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
	response.Success(ctx, gin.H{"song": result}, "正在播放")
}

//收藏歌曲
func Favor(ctx *gin.Context) {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//获取歌曲id
	id := ctx.Param("id")
	songId, _ := strconv.Atoi(id)

	var fl model.FavorList
	db_result1 := db.PGEngine.Table("favor_lists").Where(&model.FavorList{UserId: u.ID, SongId: songId}).First(&fl)

	//如果没有被收藏
	if db_result1.RecordNotFound() {
		//查询最近列表中是否存在
		var rp model.RecentList
		db_result := db.PGEngine.Table("recent_lists").Where("song_id = ?", songId).First(&rp)

		//若找不到
		if db_result.RecordNotFound() {
			var song model.Song
			song_result := db.PGEngine.Where("id = ?", songId).First(&song)

			if song_result.RecordNotFound() {
				response.Fail(ctx, nil, "系统暂无此歌曲")
			} else {
				newFavor := model.FavorList{
					UserId:    u.ID,
					SongName:  song.Name,
					SongId:    songId,
					PlayCount: 0,
				}
				db.PGEngine.Table("favor_lists").Create(&newFavor)
				//喂推荐
				fm := mygorse.ToFeedbackModel(newFavor)
				mygorse.Feedback(fm)

				//返回结果
				response.Success(ctx, nil, "收藏成功")
			}
		} else {
			newFavor := model.FavorList{
				UserId:    u.ID,
				SongName:  rp.SongName,
				SongId:    songId,
				PlayCount: rp.PlayCount,
			}
			db.PGEngine.Table("favor_lists").Create(&newFavor)
			//喂推荐
			fm := mygorse.ToFeedbackModel(newFavor)
			mygorse.Feedback(fm)

			//返回结果
			response.Success(ctx, nil, "收藏成功")
		}
	} else { //如果收藏了
		db.PGEngine.Table("favor_lists").Where("song_id = ?", songId).Delete(&fl)

		//返回结果
		response.Success(ctx, nil, "取消收藏成功")
	}
}

//获取收藏列表
func GetFavorList(ctx *gin.Context) {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//根据用户ID查询收藏列表
	var fl []model.FavorList
	db.PGEngine.Table("favor_lists").Where("user_id = ?", u.ID).Find(&fl)

	//返回结果
	var results []Result
	for key, _ := range fl {
		song := new(model.Song)
		db.PGEngine.Table("songs").Where("id = ?", fl[key].SongId).First(&song)
		artist := new(model.Artist)
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
	response.Success(ctx, gin.H{"list": results}, "收藏列表")
}

//获取最近播放列表
func GetRecentList(ctx *gin.Context) {
	//获取userInfo
	user, _ := ctx.Get("user")
	u := user.(model.User)

	//根据用户ID查询收藏列表
	var fl []model.FavorList
	db.PGEngine.Table("recent_lists").Where("user_id = ?", u.ID).Find(&fl)

	//返回结果
	var results []Result
	for key, _ := range fl {
		song := new(model.Song)
		db.PGEngine.Table("songs").Where("id = ?", fl[key].SongId).First(&song)
		artist := new(model.Artist)
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
	response.Success(ctx, gin.H{"list": results}, "最近播放列表")
}
