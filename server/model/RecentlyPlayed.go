package model

import "github.com/jinzhu/gorm"

type RecentlyPlayed struct {
	gorm.Model
	UserId    uint
	SongName  string
	SongId    int
	PlayCount int
}

