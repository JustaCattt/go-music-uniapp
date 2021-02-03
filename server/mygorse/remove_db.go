package mygorse

import (
	"log"
	"os"
)

func RemoveGorseDb() {
	//删除文件
	err := os.Remove("./mygorse/db/gorse.db")
	if err != nil {
		log.Println(err)
	}
}
