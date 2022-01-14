package rdx

import (
	"sync"

	"github.com/go-redis/redis"

	"jw.lib/conf"
)

var pool sync.Map

const defaultRedisAppName = "defaultRedisAppName"
const DefaultRedisAddr = "150.158.7.96:16379"

func Register(conn *conf.Connector, confFunc ...func()) {
	cli := redis.NewClient(&redis.Options{
		Addr:     conn.GetAddr(),
		Password: "",
	})

	pool.LoadOrStore(defaultRedisAppName, conn.GetAppName())
	pool.LoadOrStore(conn.GetAppName(), cli)
}

func GetRdxOperator(rdxName ...string) *redis.Client {
	var name interface{}
	if len(rdxName) == 0 {
		name, _ = pool.Load(defaultRedisAppName)
	} else {
		name, _ = pool.Load(rdxName[0])
	}

	res, _ := pool.Load(name)

	return res.(*redis.Client)
}
