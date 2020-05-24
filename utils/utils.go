package utils

import (
	"math/rand"
	"time"
)

// 生成随机字符串
func RandomString(num int) string {
	var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	b := make([]rune, num)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
