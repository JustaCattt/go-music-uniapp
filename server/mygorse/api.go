package mygorse

import (
	"fmt"
	"go-music-uniapp/server/myhttp"
)

const RECOMMENDER_API = "http://127.0.0.1:8080" //推荐服务

func GetRecommends(ID uint, NUM int) (rs []Recommend) {
	//拼接推荐服务的请求url，并解析到rs中
	url := fmt.Sprintf("%s/recommends/%d?number=%d", RECOMMENDER_API, ID, NUM)
	myhttp.HttpGetParser(url, &rs)
	return rs
}

func Feedback(m FeedbackModel) {
	url := fmt.Sprintf("%s/feedback", RECOMMENDER_API)
	data := fmt.Sprintf("[{\"UserId\":\"%d\",\"ItemId\":\"%d\",\"Feedback\":1}]", m.UserId, m.ItemId)
	fmt.Println(data)
	myhttp.HttpPut(url, data)
}
