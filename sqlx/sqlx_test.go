package sqlx

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"

	"jw.lib/conf"
)

func Test_pg(t *testing.T) {
	Register(DefaultSqlDriver, conf.APP_PG_ADDR.Value(DefaultSqlAddr))

	stmt, err := GetSqlOperator().Prepare("select * from test")
	if err != nil {
		log.Print(err)
		return
	}

	r, err := stmt.Query()
	if err != nil {
		log.Print(err)
		return
	}

	var list []string
	for r.Next() {
		id := ""
		err = r.Scan(&id)
		if err != nil {
			log.Print(err)
			return
		}
		list = append(list, id)
	}

	fmt.Println(list)
}
