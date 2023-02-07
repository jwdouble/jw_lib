package rdx

import (
	"context"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"

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

func Register(m map[string]string, rdxName ...string) {
	cli := redis.NewClient(&redis.Options{
		Addr:     m["addr"],
		Password: m["password"],
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		log.Println("rdx.Register: redis.Ping failed: ", err)
		return
	}

	if len(rdxName) == 0 {
		pool.LoadOrStore("redis", cli)
	} else {
		pool.LoadOrStore(rdxName[0], cli)
	}
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

	return nil
}
