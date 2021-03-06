package sqlx

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
)

func Test_pgGet(t *testing.T) {

	Register(Driver, PGConfigMap)

	stmt, err := GetSqlOperator().Prepare("select * from test_text")
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
		text := ""
		err = r.Scan(&id, &text)
		if err != nil {
			log.Print(err)
			return
		}
		list = append(list, text)
	}

	fmt.Println(list)
}

func Test_pgAdd(t *testing.T) {
	Register(Driver, PGConfigMap)
	stmt, err := GetSqlOperator().Prepare(`insert into test_text(t,id) values ($1, $2)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec("777", 2)
	if err != nil {
		log.Fatal().Err(err)
	}
}
