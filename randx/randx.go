package randx

import (
	"math/rand"
	"time"
	"unsafe"
)

// NewString 生成随机字符串
func NewString(n int, seed ...int64) string {
	if n > 1024 {
		return "length over 1024"
	}
	source := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	l := len(source)

	if len(seed) > 0 {
		var count int64
		for _, v := range seed {
			count += v
		}

		rand.Seed(count)
	} else {
		rand.Seed(time.Now().Unix())
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = source[rand.Intn(l)]
	}

	return *(*string)(unsafe.Pointer(&buf))
}

func NewInt(n int, seed ...int64) int {
	if len(seed) > 0 {
		var count int64
		for _, v := range seed {
			count += v
		}

		r := rand.New(rand.NewSource(count))
		return r.Intn(n)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(n)
}
