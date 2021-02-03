package dataGetter

import (
	"fmt"
	"go-music-uniapp/server/model"
	"go-music-uniapp/server/myhttp"
)

//网易云开源接口：http://musicapi.leanapp.cn/
//这里是有人将这个项目上传到服务器上了，所以我直接使用这个api了
const API = "https://autumnfish.cn"

type ArtistSlice struct {
	Artists []model.Artist
}

type SongSlice struct {
	Songs []model.Song
}

func GetArtists() []model.Artist {
	url := fmt.Sprintf("%s/artist/list", API)
	var s ArtistSlice
	myhttp.HttpGetParser(url, &s)
	return s.Artists
}

func GetAllSongsByArtistId(artistId int) []model.Song {
	url := fmt.Sprintf("%s/artist/songs?id=%d", API, artistId)
	var s SongSlice
	myhttp.HttpGetParser(url, &s)
	return s.Songs
}

func GetTop50SongsByArtistId(artistId int) []model.Song {
	url := fmt.Sprintf("%s/artist/top/song?id=%d", API, artistId)
	var s SongSlice
	myhttp.HttpGetParser(url, &s)
	return s.Songs
}
