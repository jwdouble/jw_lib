package logx

import (
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

var commonLog = Log{zerolog.New(NewIoWriter()).With().Timestamp().Logger()}
var errLog = Log{zerolog.New(NewIoWriter()).With().Caller().Timestamp().Logger()}

func init() {
	zerolog.LevelFieldName = "l"
	zerolog.TimestampFieldName = "t"
	zerolog.TimeFieldFormat = timex.DateTimeFormat
	zerolog.MessageFieldName = "msg"
	zerolog.CallerSkipFrameCount = 4
}

// 这里如果返回的是非指针结构体，KV失效。探究
func getLogger() *Log {
	return &commonLog
}

func getErrLogger() *Log {
	return &errLog
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

func (l *Log) KV(key string, value interface{}) {
	l.Logger = l.Logger.With().Interface(key, value).Logger()
}
