package dataGetter

import (
	"encoding/json"
	"fmt"
	"go-music-uniapp/server/model"
	"io/ioutil"
	"net/http"
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
	HttpGetParser(url, &s)
	return s.Artists
}

func GetAllSongsByArtistId(artistId int) []model.Song {
	url := fmt.Sprintf("%s/artist/songs?id=%d", API, artistId)
	var s SongSlice
	HttpGetParser(url, &s)
	return s.Songs
}

func GetTop50SongsByArtistId(artistId int) []model.Song {
	url := fmt.Sprintf("%s/artist/top/song?id=%d", API, artistId)
	var s SongSlice
	HttpGetParser(url, &s)
	return s.Songs
}

func HttpGetParser(url string, v interface{}) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic("http get error")
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, v)
}
