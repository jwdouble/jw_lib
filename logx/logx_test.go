package logx

import (
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"jw.lib/timex"
)

func Test_zerolog(t *testing.T) {
	zerolog.TimeFieldFormat = timex.DateTimeFormat

	log.Print("hello world")
	log.Info().Msg("hello world")

	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send() // Send is equivalent to calling Msg("")

	err := errors.New("seems we have an error here")
	log.Error().Err(err).Msg("haha")
}

func Test_mylog(t *testing.T) {
	err := errors.New("this is error")
	Errorf(err, "this is error %s", "test")
}

func Test_redis(t *testing.T) {
	//Info("INFO TEST")
	//
	//sc := rdx.GetRdxOperator().RPop("logx")
	//r, err := sc.Result()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(r)
}
