package sqlx

import (
	"database/sql"
	"jw.lib/toolx"
	"sync"

	_ "github.com/lib/pq"
)

var pool sync.Map

type SqlMod struct {
	db *sql.DB
}

// TODO: pg的注册还可以优化

const (
	Driver = "postgres"
)

func Register(driver string, cm map[string]string) {
	addr := toolx.RenderString("host=${host} user=${user} password=${password} dbname=${dbname} port=${port} ", cm)
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
		val, _ = pool.Load(Driver)
	} else {
		val, _ = pool.Load(dbName[0])
	}

	res, ok := val.(*SqlMod)
	if ok {
		return res.db
	}

	panic("redis not register")
}
