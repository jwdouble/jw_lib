package conf

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvVar_GetEnv(t *testing.T) {
	t.Log(Get("lib.logx.redis"))
}

func Test_envVar(t *testing.T) {
	//os.Setenv("GO_COMMON_PASSWORD", "jw")
	fmt.Println(os.Getenv("GO_COMMON_PASSWORD"))
}
