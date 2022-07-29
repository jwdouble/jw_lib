package rdx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	Register(RedisConfigMap)

	GetRdxOperator().Set(context.Background(), "test", "1", time.Minute)

	cmd := GetRdxOperator().Get(context.Background(), "test")
	fmt.Println(cmd.Val())
}
