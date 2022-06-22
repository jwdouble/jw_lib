package randx

import (
	"math/rand"
	"time"
	"unsafe"
)

// NewString 生成随机字符串
func NewString(n int) string {
	if n > 1024 {
		return "length over 1024"
	}
	source := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	l := len(source)

	rand.Seed(time.Now().Unix())
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = source[rand.Intn(l)]
	}

	return *(*string)(unsafe.Pointer(&buf))
}

func NewInt(n int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(n)
}
