package logx

import (
	"fmt"
	"testing"

	"jw.lib/rdx"
)

func Test_log(t *testing.T) {
	Info("this is %s code, no.%d", []any{"jw", 1})
}

func Test_redis(t *testing.T) {
	Info("INFO TEST")

	sc := rdx.GetRdxOperator().RPop("logx")
	r, err := sc.Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)
}
