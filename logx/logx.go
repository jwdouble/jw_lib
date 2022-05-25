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
)

var pool = sync.Pool{
	New: func() interface{} {
		return &Logger{Std: "redis"}
	},
}

// 加个调试堆栈

type Logger struct {
	CreateAt time.Time     `json:"createAt"`
	Level    zerolog.Level `json:"Alevel,omitempty"`
	Position string        `json:"position,omitempty"`
	FuncName string        `json:"funcName"`
	Content  string        `json:"content,omitempty"`
	Std      string        `json:"std,omitempty"`
}

func (l *Logger) Write() {
	switch l.Std {
	case "postgres":
		logToPg(l)
	case "redis":
		logToRedis(l)
	default:
		logToPg(l)
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
	stmt, err := cli.Prepare("insert into service (create_time, level, position, content) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(l.CreateAt, l.Level, l.Position, l.Content)

	if err != nil {
		panic(err.Error())
	}
}

func logToRedis(l *Logger) {
	onceRedis.Do(func() {
		rdx.Register(conf.APP_REDIS_ADDR.Value(rdx.DefaultRedisAddr), "")
	})

	cli := rdx.GetRdxOperator()
	// 用redis的有序集合 存贮log， jw sys每10s读一个集合
	// 优化: 用列表，添加新日志时间常数级别。
	buf, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
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
	ptr, file, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(ptr)
	l := pool.Get().(*Logger)
	l.FuncName = f.Name()
	l.Position = file + ":" + strconv.Itoa(line)
	l.CreateAt = time.Now()
	l.Level = level
	switch err.(type) {
	case string:
		l.Content = err.(string)
	case error:
		l.Content = err.(error).Error()
	}

	l.Write()
	pool.Put(l)
}
