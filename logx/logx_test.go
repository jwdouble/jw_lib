package logx

import (
	"fmt"
	"jw.lib/rdx"
	"testing"
)

func Test_logx(t *testing.T) {
	Info("INFO TEST")

	sc := rdx.GetRdxOperator().RPop("logx")
	r, err := sc.Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)
}
