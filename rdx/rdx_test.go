package rdx

import (
	"fmt"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	Register(RedisConfigMap)

	GetRdxOperator().Set("test", "1", time.Minute)

	cmd := GetRdxOperator().Get("test")
	fmt.Println(cmd.Val())
}
