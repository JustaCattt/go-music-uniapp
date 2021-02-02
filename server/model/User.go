package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"type:varchar(20);not null" json:"username"`
	Telephone string    `gorm:"type:varchar(110);not null;unique" json:"telephone"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	FavorList FavorList `gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	RecentlyPlayed RecentlyPlayed `gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
}
