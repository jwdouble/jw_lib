package sqlx

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

// TODO config

var pool sync.Map

type SqlMod struct {
	db *sql.DB
}

const (
	DefaultSqlDriver = "postgres"
	DefaultSqlAddr   = "host=150.158.7.96 user=postgres password=p1ssw0rd dbname=jwdouble port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func Register(driver, addr string) {
	db, err := sql.Open(driver, addr)
	if err != nil {
		panic(err)
	}

	pg := &SqlMod{db: db}
	pool.LoadOrStore(driver, pg)
}

func GetSqlOperator(dbName ...string) *sql.DB {
	var val interface{}

	if len(dbName) == 0 {
		val, _ = pool.Load(DefaultSqlDriver)
	} else {
		val, _ = pool.Load(dbName[0])
	}

	res, ok := val.(*SqlMod)
	if ok {
		return res.db
	}

	panic("redis not register")
}
