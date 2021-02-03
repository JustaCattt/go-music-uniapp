package mygorse

import "go-music-uniapp/server/model"

type Recommend struct {
	ItemId     string
	Popularity int
	Timestamp  string
	Score      int
}

type FeedbackModel struct {
	UserId uint
	ItemId int
}

//FavorListToFeedbackModel
func ToFeedbackModel(fl model.FavorList) FeedbackModel {
	return FeedbackModel{
		UserId: fl.UserId,
		ItemId: fl.SongId,
	}
}
