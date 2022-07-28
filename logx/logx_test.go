package logx

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"jw.lib/rdx"
)

func Test_zerolog(t *testing.T) {
	Infof("INFO TEST1")

	KV("app", "jw-lib")

	Infof("INFO TEST6")
}

func Test_mylog(t *testing.T) {
	err := errors.New("this is error")
	Errorf(err, "this is error %s", "test")
}

func Test_redis(t *testing.T) {
	//KV("app", "test")
	//
	//Infof("INFO TEST")
	sc := rdx.GetRdxOperator().LPop("logx")

	r := sc.Val()
	fmt.Println("log from redis -->", r)

	r = sc.String()
	fmt.Println("log from redis -->", r)

	r, err := sc.Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("log from redis -->", r)
}

func Test_xx(t *testing.T) {
	l := GetE()

	l.Debug().Caller().Msg("test")
}

func GetE() zerolog.Logger {
	l := log.With().Caller().Logger()
	return l
}

func Test_err(t *testing.T) {
	err := errors.New("haha")
	Errorf(err, "this is error %s", "test")
}
