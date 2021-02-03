package model

import "github.com/jinzhu/gorm"

type RecentList struct {
	gorm.Model
	UserId    uint
	SongName  string
	SongId    int
	PlayCount int
}
