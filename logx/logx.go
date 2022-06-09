package logx

import (
	"encoding/json"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"jw.lib/conf"
	"jw.lib/rdx"
	"jw.lib/sqlx"
	"jw.lib/timex"
)

var pool = sync.Pool{
	New: func() interface{} {
		return &Logger{std: conf.Get("lib.logx.driver")}
	},
}

type Logger struct {
	Ts     string        `json:"ts"`
	Level  zerolog.Level `json:"level,omitempty"`
	Caller string        `json:"caller,omitempty"`
	Msg    string        `json:"msg,omitempty"`
	std    string
}

func (l *Logger) Write() {
	switch l.std {
	case "postgres":
		logToPg(l)
	case "redis":
		logToRedis(l)
	default:
		logToRedis(l)
	}
}

var (
	oncePg    sync.Once
	onceRedis sync.Once
)

func logToPg(l *Logger) {
	oncePg.Do(func() {
		sqlx.Register(sqlx.DefaultSqlDriver, conf.APP_PG_ADDR.Value(sqlx.DefaultSqlAddr))
	})

	cli := sqlx.GetSqlOperator()
	// TODO: auto create table
	stmt, err := cli.Prepare(`insert into logx (create_time, level, position, content) values($1, $2, $3, $4)`)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(l.Ts, l.Level, l.Caller, l.Msg)

	if err != nil {
		panic(err.Error())
	}
}

func logToRedis(l *Logger) {
	onceRedis.Do(func() {
		rdx.Register(conf.APP_REDIS_ADDR.Value(rdx.DefaultRedisAddr), "jw")
	})

	cli := rdx.GetRdxOperator()
	// 用redis的有序集合 存贮log， jw sys每10s读一个集合
	// 优化: 用列表，添加新日志时间常数级别。
	buf, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	// list 列表 RPush
	cli.RPush("logx", string(buf))
}

func Info(err interface{}) {
	newLogger(err, zerolog.InfoLevel)
}

func Debug(err interface{}) {
	newLogger(err, zerolog.DebugLevel)
}

func Warn(err interface{}) {
	newLogger(err, zerolog.WarnLevel)
}

func Error(err interface{}) {
	newLogger(err, zerolog.ErrorLevel)
}

func newLogger(err interface{}, level zerolog.Level) {
	l := pool.Get().(*Logger)

	_, file, line, _ := runtime.Caller(2) // prt, file, line
	//f := runtime.FuncForPC(ptr)
	//l.FuncName = f.Name()
	l.Caller = file + ": " + strconv.Itoa(line)
	l.Ts = time.Now().Format(timex.DateTimeFormat)
	l.Level = level
	switch err.(type) {
	case string:
		l.Msg = err.(string)
	case error:
		l.Msg = err.(error).Error()
	}

	l.Write()
	pool.Put(l)
}
