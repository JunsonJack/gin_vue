package util

import (
	"math/rand"
	"time"
)

func RandomStr(n int) string {
	var str = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKL")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = str[rand.Intn(len(str))]
	}
	return string(result)
}

