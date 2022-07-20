package sqlx

import (
	"database/sql"
	"sync"

	"jw.lib/conf"
	"jw.lib/toolx"

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

var PGConfigMap = map[string]string{
	"host":     "150.158.7.96",
	"user":     "postgres",
	"password": conf.COMMON_PASSWORD.Value("xxx"),
	"dbname":   "jwdouble",
	"port":     "5432",
	"suffix":   "sslmode=disable TimeZone=Asia/Shanghai",
}

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
