package dataGetter

import (
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/model"
)

var ch = make(chan model.Artist)

func InsertArtists() {
	artists := GetArtists()
	for _, artist := range artists {
		ch <- artist
		db.PGEngine.Create(&artist)
	}
}

func InsertSongs() {
	for artist := range ch {
		songs := GetAllSongsByArtistId(artist.Id)
		for _, song := range songs {
			song.Artist = artist
			song.ArtistId = artist.Id
			db.PGEngine.Create(&song)
		}
	}
}

func MapArtistToSongs() map[string][]model.Song {
	artists := GetArtists()
	m := make(map[string][]model.Song, len(artists))
	for _, artist := range artists {
		songs := GetAllSongsByArtistId(artist.Id)
		m[artist.Name] = songs
	}
	return m
}

func InitData() {
	go InsertArtists()
	go InsertSongs()
}
