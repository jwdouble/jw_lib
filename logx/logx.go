package logx

import (
	"os"

	"github.com/rs/zerolog"
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
	l.Error().Err(err).Msgf(str, arg...)
}

func (l *Log) Fatalf(str string, arg ...interface{}) {
	l.Fatal().Msgf(str, arg...)
}

//func (l *Logger) write() {
//	switch l.std {
//	case "postgres":
//		logToPg(l)
//	case "redis":
//		logToRedis(l)
//	default:
//		fmt.Println(l)
//	}
//}
//
//var (
//	oncePg    sync.Once
//	onceRedis sync.Once
//)
//
//func logToPg(l *Logger) {
//	oncePg.Do(func() {
//		sqlx.Register(sqlx.Driver, sqlx.PGConfigMap)
//	})
//
//	cli := sqlx.GetSqlOperator()
//	// TODO: auto create table
//	stmt, err := cli.Prepare(`insert into logx (create_time, level, position, content) values($1, $2, $3, $4)`)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	_, err = stmt.Exec(l.Ts, l.Level, l.Caller, l.Msg)
//
//	if err != nil {
//		panic(err.Error())
//	}
//}
//
//func logToRedis(l *Logger) {
//	onceRedis.Do(func() {
//		rdx.Register(rdx.RedisConfigMap)
//	})
//
//	cli := rdx.GetRdxOperator()
//	// 用redis的有序集合 存贮log， jw sys每10s读一个集合
//	// 优化: 用列表，添加新日志时间常数级别。
//	buf, err := json.Marshal(l)
//	if err != nil {
//		panic(err)
//	}
//	// list 列表 RPush
//	cli.RPush("logx", string(buf))
//}
//
//func Info(format string, arg ...any) {
//	newLogger(fmt.Sprintf(format, arg), zerolog.InfoLevel, arg)
//}
//
//func Debug(format string, withPos bool, arg ...any) {
//	if withPos {
//		newLogger(fmt.Sprintf(format, arg), zerolog.DebugLevel, DisplayPos)
//	} else {
//		newLogger(fmt.Sprintf(format, arg), zerolog.DebugLevel)
//	}
//
//}
//
//func Warn(format string, arg ...any) {
//	newLogger(fmt.Sprintf(format, arg), zerolog.WarnLevel, arg)
//}
//
//func Error(format string, arg ...any) {
//	for n := range arg {
//		if e, ok := arg[n].(error); ok {
//			arg[n] = e.Error()
//		}
//	}
//
//	newLogger(fmt.Sprintf(format, arg), zerolog.ErrorLevel, DisplayPos)
//}
//
//func newLogger(msg interface{}, level zerolog.Level, optional ...any) {
//	l := pool.Get().(*Logger)
//
//	m := map[string]bool{}
//	for _, v := range optional {
//		m[v.(string)] = true
//	}
//
//	ptr, file, line, _ := runtime.Caller(2)
//	l.Caller = runtime.FuncForPC(ptr).Name()
//
//	if m[DisplayPos] {
//		l.Pos = file + ": " + strconv.Itoa(line)
//	}
//
//	l.Ts = time.Now().Format(timex.DateTimeFormat)
//
//	l.Level = level
//
//	switch msg.(type) {
//	case string:
//		l.Msg = msg.(string)
//	case error:
//		l.Msg = msg.(error).Error()
//	}
//
//	l.write()
//	pool.Put(l)
//}
