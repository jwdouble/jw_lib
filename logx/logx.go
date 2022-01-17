package logx

import (
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
		//return &Logger{std: conf.GetYaml("app.engine.log")}
		return &Logger{std: "redis"}
	},
}

// 加个调试堆栈

type Logger struct {
	createAt time.Time
	level    zerolog.Level
	// todo 返回调用的方法名 --最好能返回在哪一行出的错
	position string
	content  string
	std      string
}

func (l *Logger) Write() {
	switch l.std {
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
		sqlx.Register(conf.AppPgConn.Value(sqlx.DefaultPgAddr))
	})

	cli := sqlx.GetSqlOperator()
	stmt, err := cli.Prepare("insert into service (create_time, level, position, content) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(l.createAt, l.level, l.position, l.content)

	if err != nil {
		panic(err.Error())
	}
}

func logToRedis(l *Logger) {
	onceRedis.Do(func() {
		rdx.Register(conf.AppRedisConn.Value(rdx.DefaultRedisAddr))
	})

	cli := rdx.GetRdxOperator()
	cli.Set("logx-"+l.createAt.Format(timex.DateTimeFormat)+"-"+l.level.String(), l.content, time.Hour*24*8)
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
	l.createAt = time.Now()
	l.level = level
	switch err.(type) {
	case string:
		l.content = err.(string)
	case error:
		l.content = err.(error).Error()
	}

	l.Write()
	pool.Put(l)
}
