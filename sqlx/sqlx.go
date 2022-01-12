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

const DefaultPgAddr = "host=150.158.7.96 user=postgres password=p1ssw0rd " +
	"dbname=logs port=25432 sslmode=disable TimeZone=Asia/Shanghai"

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
	var name interface{}
	if len(dbName) == 0 {
		name, _ = pool.Load(DefaultPggAppName)
	} else {
		name, _ = pool.Load(dbName[0])
	}

	res, _ := pool.Load(name)

	return res.(*SqlMod).db
}
