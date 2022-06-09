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

const pos = "pos"

var pool = sync.Pool{
	New: func() interface{} {
		return &Logger{std: conf.Get("lib.logx.driver")}
	},
}

type Logger struct {
	// 时间
	Ts    string        `json:"ts"`
	Level zerolog.Level `json:"level"`
	// 函数名
	Caller string `json:"caller"`
	Msg    string `json:"msg"`
	// 代码位置 可选 默认关闭原因：信息太长影响美观
	Pos string `json:"pos,omitempty"`
	std string
}

func (l *Logger) write() {
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

func Info(info interface{}, optional ...string) {
	newLogger(info, zerolog.InfoLevel, optional...)
}

func Debug(debug interface{}, optional ...string) {
	newLogger(debug, zerolog.DebugLevel, optional...)
}

func Warn(warn interface{}, optional ...string) {
	newLogger(warn, zerolog.WarnLevel, optional...)
}

func Error(err interface{}, optional ...string) {
	newLogger(err, zerolog.ErrorLevel, optional...)
}

func newLogger(msg interface{}, level zerolog.Level, optional ...string) {
	l := pool.Get().(*Logger)

	m := map[string]bool{}
	for _, v := range optional {
		m[v] = true
	}

	ptr, file, line, _ := runtime.Caller(2)
	l.Caller = runtime.FuncForPC(ptr).Name()

	if m[pos] {
		l.Pos = file + ": " + strconv.Itoa(line)
	}

	l.Ts = time.Now().Format(timex.DateTimeFormat)

	l.Level = level

	switch msg.(type) {
	case string:
		l.Msg = msg.(string)
	case error:
		l.Msg = msg.(error).Error()
	}

	l.write()
	pool.Put(l)
}
