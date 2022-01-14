package sqlx

import (
	"database/sql"
	"sync"

	"jw.lib/conf"

	_ "github.com/lib/pq"
)

// TODO config

var pool sync.Map

type SqlMod struct {
	db *sql.DB
}

const DefaultPgAppName = "defaultPgAppName"
const DefaultPgAddr = "host=150.158.7.96 user=postgres password=p1ssw0rd " +
	"dbname=logs port=25432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	Register(conf.AppPgConn.Value(DefaultPgAddr))
}

func Register(conn *conf.Connector, confFunc ...func(db *sql.DB)) {
	db, err := sql.Open(conn.GetDriverName(), conn.GetAddr())
	if err != nil {
		panic(err)
	}

	for _, f := range confFunc {
		f(db)
	}

	pg := &SqlMod{db: db}
	pool.LoadOrStore(DefaultPgAppName, conn.GetAppName())
	pool.LoadOrStore(conn.GetAppName(), pg)
}

func GetSqlOperator(dbName ...string) *sql.DB {
	var name interface{}
	if len(dbName) == 0 {
		name, _ = pool.Load(DefaultPgAppName)
	} else {
		name, _ = pool.Load(dbName[0])
	}

	res, _ := pool.Load(name)

	return res.(*SqlMod).db
}
