package model

type Song struct {
	Id       int    `gorm:"primary_key;"`
	Name     string `gorm:"size:255;not null;"`
	Artist   Artist `gorm:"ForeignKey:ArtistId;AssociationForeignKey:Id"`
	ArtistId int
}
