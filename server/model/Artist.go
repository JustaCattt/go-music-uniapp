package model

type Artist struct {
	Id    int    `gorm:"primary_key;"`
	Name  string `gorm:"type:varchar(128);not null;"`
	Songs []Song `gorm:"ForeignKey:ArtistId;AssociationForeignKey:Id"`
}
