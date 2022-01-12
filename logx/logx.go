package logx

import (
	"sync"
	"time"

	"github.com/rs/zerolog"

	"jw.lib/conf"
	"jw.lib/sqlx"
)

var pool = sync.Pool{
	New: func() interface{} {
		return &Logger{std: conf.GetYaml("app.engine.log")}
	},
}

// 加个调试堆栈

type Logger struct {
	level zerolog.Level
	// 返回调用的方法名 --最好能返回在哪一行出的错
	position string
	content  string
	std      string
}

func (l *Logger) Write() {
	switch l.std {
	case "postgres":
		logToPg(l)
	case "redis":
	default:
		logToPg(l)
	}
}

var (
	oncePg    sync.Once
	onceRedis sync.Once
)

func logToPg(l *Logger) {
	//oncePg.Do(func() {
	sqlx.Register(conf.AppPgConn.Value(sqlx.DefaultPgAddr), nil)
	//})

	client := sqlx.GetSqlOperator()
	stmt, err := client.Prepare("insert into service (create_time, level, position, content) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(time.Now(), l.level, l.position, l.content)

	if err != nil {
		panic(err.Error())
	}
}

func Info(s string) {
	l := pool.Get().(*Logger)
	l.content = s
	l.level = zerolog.InfoLevel
	l.Write()
	pool.Put(l)
}

func Debug(s string) {
	l := pool.Get().(*Logger)
	l.content = s
	l.level = zerolog.DebugLevel
	l.Write()
	pool.Put(l)
}

func Warn(s string) {
	l := pool.Get().(*Logger)
	l.content = s
	l.level = zerolog.WarnLevel
	l.Write()
	pool.Put(l)
}

func Error(s string) {
	l := pool.Get().(*Logger)
	l.content = s
	l.level = zerolog.ErrorLevel
	l.Write()
	pool.Put(l)
}
