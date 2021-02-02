package data

import (
	"fmt"
	"go-music-uniapp/server/db"
	"go-music-uniapp/server/util"
	"os"
)

type Result struct {
	UserId    uint
	SongId    int
	PlayCount int
}

func DbToCsv() {
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
	workDir, _ := os.Getwd() //获取工作目录
	dir := fmt.Sprintf("%s/data/data.csv", workDir)
	util.CsvWriter(dir, data)
}
