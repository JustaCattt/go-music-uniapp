package db_to_csv

import (
	"fmt"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/util"
)

type Result struct {
	UserId    uint
	SongId    int
	PlayCount int
}

type Item struct {
	Id int
}

func InitData() {
	util.RemoveCsvData("./data/data.csv")

	var results []Result
	db.PGEngine.Table("favor_lists").Find(&results)
	var data [][]string
	for index, _ := range results {
		fmt.Println(results[index])
		result := results[index]
		uIdStr := fmt.Sprintf("%d", result.UserId)
		sIdStr := fmt.Sprintf("%d", result.SongId)
		pcStr := fmt.Sprintf("%d", result.PlayCount)
		data = append(data, []string{uIdStr, sIdStr, pcStr})
	}

	util.CsvWriter("./data/data.csv", data)
}

func InitItem() {
	util.RemoveCsvData("./data/item.csv")

	var items []Item
	db.PGEngine.Table("songs").Find(&items)
	var data [][]string
	for index, _ := range items {
		fmt.Println(items[index])
		item := items[index]
		sIdStr := fmt.Sprintf("%d", item.Id)
		data = append(data, []string{sIdStr})
	}

	util.CsvWriter("./data/item.csv", data)
}
