package rdx

import (
	"fmt"
	"jw.lib/conf"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	Register(conf.APP_REDIS_ADDR.Value(DefaultRedisAddr), RedisPwd)

	GetRdxOperator().Set("test", "1", time.Minute)

	cmd := GetRdxOperator().Get("test")
	fmt.Println(cmd.Val())
}
