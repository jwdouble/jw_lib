package conf

import (
	"github.com/spf13/viper"
	"golang.org/x/sys/unix"
	"io/fs"
	"log"
)

var (
	filePath = "./config.yaml"
)

var vip *viper.Viper

func init() {
	vip = viper.New()
	vip.SetConfigFile(filePath)
	err := vip.ReadInConfig()
	if err != nil {
		e := err.(*fs.PathError)
		if e.Err == unix.ENOENT {
			log.Fatal("conf file not found")
		} else {
			panic(err)
		}
	}
}

func Get(key string) string {
	val := vip.Get(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
