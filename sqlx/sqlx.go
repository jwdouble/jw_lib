package sqlx

import (
	"database/sql"
	"sync"

	"jw.lib/conf"
)

// TODO config

var pool sync.Map

type SqlMod struct {
	db *sql.DB
}

const DefaultPgAddr = "host=150.158.7.96 user=postgres password=p1ssw0rd " +
	"dbname=sys_log port=25432 sslmode=disable TimeZone=Asia/Shanghai statement_cache_mode=describe"

const DefaultPggAppName = "defaultPgAppName"

func init() {
	Register(conf.AppPgConn.Value(DefaultPgAddr), nil)
}

func Register(c *conf.Connector, confFunc func(db *sql.DB)) {
	db, err := sql.Open(c.GetDriverName(), c.GetDSN())
	if err != nil {
		panic(err)
	}

	if confFunc != nil {
		confFunc(db)
	}

	pg := &SqlMod{db: db}
	pool.LoadOrStore(DefaultPggAppName, c.GetAppName())
	pool.LoadOrStore(c.GetAppName(), pg)
}

func GetSqlOperator(dbName ...string) *sql.DB {
	name := ""
	if len(dbName) == 0 {
		name = DefaultPggAppName
	} else {
		name = dbName[0]
	}

	ans, _ := pool.Load(name)

	return ans.(*SqlMod).db
}
