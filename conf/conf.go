package conf

import (
	"os"
)

func init() {
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		logx
	}
}
