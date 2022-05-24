package conf

import (
	"fmt"
	"os"
	"testing"
)

func Test_envVar(t *testing.T) {
	fmt.Println(APP_PORT.Value("haha"))
	os.Setenv("K8S_APP_PORT", "xxx")
	fmt.Println(APP_PORT.Value("kkk"))
}
