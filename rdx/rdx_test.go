package rdx

import (
	"fmt"
	"testing"
	"time"

	"jw.lib/conf"
)

func Test_redis(t *testing.T) {
	Register(conf.APP_REDIS_ADDR.Value(DefaultRedisAddr), "")

	GetRdxOperator().Set("test", "1", time.Minute)

	cmd := GetRdxOperator().Get("test")
	fmt.Println(cmd.Val())
}
