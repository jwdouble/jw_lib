package logx

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/go-redis/redis/v8"

	"jw.lib/rdx"
)

type std struct {
	std io.Writer
	rc  *redis.Client
}

var redisCtxBackground = context.Background()

func NewIoWriter() *std {
	rdx.Register(rdx.RedisConfigMap)
	return &std{
		std: os.Stderr,
		rc:  rdx.GetRdxOperator(),
	}
}

func (s *std) Write(p []byte) (n int, err error) {
	n, err = s.std.Write(p)
	if err != nil {
		return
	}

	// 用redis的有序集合 存贮log， jw sys每10s读一个集合
	// 优化: 用列表，添加新日志时间常数级别。
	// list 列表 RPush

	if s.rc == nil {
		log.Println("logx.redis.Write: redis.Client is nil")
		return
	}

	ic := s.rc.RPush(redisCtxBackground, "logx", string(p))
	if ic.Err() != nil { // TODO: redis写入失败就把redis失败信息通过标准输出打印？ 该处理是否合理
		return s.std.Write([]byte("logx.redis.Write: " + ic.Err().Error() + "\n"))
	}

	return
}
