package rdx

import (
	"sync"

	"github.com/go-redis/redis"

	"jw.lib/conf"
)

var pool sync.Map

const (
	DefaultRedisAddr = "150.158.7.96:6379"
)

var RedisConfigMap = map[string]string{
	"addr":     DefaultRedisAddr,
	"password": conf.COMMON_PASSWORD.Value("xxx"),
}

func Register(m map[string]string) {
	cli := redis.NewClient(&redis.Options{
		Addr:     m["addr"],
		Password: m["password"],
	})

	pool.LoadOrStore("redis", cli)
}

func GetRdxOperator(rdxName ...string) *redis.Client {
	var val interface{}

	if len(rdxName) == 0 {
		val, _ = pool.Load("redis")
	} else {
		val, _ = pool.Load(rdxName[0])
	}

	res, ok := val.(*redis.Client)
	if ok {
		return res
	}

	panic("redis not register")
}
