package logx

import (
	"os"

	"github.com/rs/zerolog"

	"jw.lib/timex"
)

//var pool = sync.Pool{
//	New: func() interface{} {
//		return &Logger{std: conf.Get("lib.logx.driver")}
//	},
//}

type Logger interface {
	Debugf(str string, arg ...interface{})
	Infof(str string, arg ...interface{})
	Warnf(str string, arg ...interface{})
	Errorf(str string, arg ...interface{})
	Fatalf(str string, arg ...interface{})
}

type Log struct {
	zerolog.Logger
}

var rootLogger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func init() {
	zerolog.LevelFieldName = "l"
	zerolog.TimestampFieldName = "t"
	zerolog.TimeFieldFormat = timex.DateTimeFormat
	zerolog.MessageFieldName = "msg"
}

func getLogger() Log {
	return Log{rootLogger}
}

func (l *Log) Debugf(str string, arg ...interface{}) {
	l.Debug().Msgf(str, arg...)
}

func (l *Log) Infof(str string, arg ...interface{}) {
	l.Info().Msgf(str, arg...)
}

func (l *Log) Warnf(str string, arg ...interface{}) {
	l.Warn().Msgf(str, arg...)
}

func (l *Log) Errorf(err error, str string, arg ...interface{}) {
	if err == nil {
		l.Error().Msgf(str, arg...)
	} else {
		l.Error().Err(err).Msgf(str, arg...)
	}
}

func (l *Log) Fatalf(str string, arg ...interface{}) {
	l.Fatal().Msgf(str, arg...)
}
