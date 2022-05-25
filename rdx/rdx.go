package rdx

import (
	"sync"

	"github.com/go-redis/redis"
)

var pool sync.Map

const (
	DefaultRedisAddr = "150.158.7.96:6379"
	RedisPwd         = "jw"
)

func Register(addr, pwd string) {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
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
