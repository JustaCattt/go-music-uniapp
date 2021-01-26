package util

import (
	"encoding/csv"
	"os"
)

func CsvWriter(filename string, data [][]string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("open file is failed, err: " + err.Error())
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF") //防止中文乱码
	w := csv.NewWriter(file)
	w.WriteAll(data)
	w.Flush() // 写文件需要flush，不然缓存满了，后面的就写不进去了，只会写一部分
}
