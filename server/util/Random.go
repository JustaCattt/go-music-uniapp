package util

import (
	"math/rand"
	"time"
)

//n位随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//0~n随机数
func RandomNum(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}
