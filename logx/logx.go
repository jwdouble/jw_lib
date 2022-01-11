package Logger

import (
	"sync"
	"time"

	"github.com/rs/zerolog"

	"jw.lib/conf"
	"jw.lib/sqlx"
)

type Log interface {
	Redirect()
	Debug()
	Warning()
	Error()
}

type Logger struct {
	Level zerolog.Level
	// 返回调用的方法名 --最好能返回在哪一行出的错
	Position string
	content  string
	std      string
}

func (l *Logger) Write() {
	switch l.std {
	case "sql":
		logToSql(l)
	case "redis":

	}
}

// Redirect 可以重定向到redis、database、other
func (l *Logger) Redirect(appName string) {
	l.std = appName
}

var (
	oncePg    sync.Once
	onceRedis sync.Once
)

func logToSql(l *Logger) {
	oncePg.Do(func() {
		sqlx.Register(conf.AppPgConn.Value(sqlx.DefaultPgAddr), nil)
	})

	client := sqlx.GetSqlOperator()
	stmt, err := client.Prepare("insert into service (create_time, level, position, content) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(time.Now, l.Level, l.Position, l.content)

	if err != nil {
		panic(err.Error())
	}
}

func (l *Logger) Info() {
	l.Level = zerolog.InfoLevel
	l.Write()
}

func (l *Logger) Debug() {
	l.Level = zerolog.DebugLevel
	l.Write()
}

func (l *Logger) Warning() {
	l.Level = zerolog.WarnLevel
	l.Write()
}

func (l *Logger) Error() {
	l.Level = zerolog.ErrorLevel
	l.Write()
}
