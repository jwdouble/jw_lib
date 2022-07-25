package conf

import (
	"io/fs"
	"log"

	"github.com/spf13/viper"
	"golang.org/x/sys/unix"
)

var vip *viper.Viper

func init() {
	vip = viper.New()

	vip.SetConfigFile(CONF_FILE_PATH)
	err := vip.ReadInConfig()
	if err != nil {
		e := err.(*fs.PathError)
		if e.Err == unix.ENOENT {
			log.Println("conf file not found: ", CONF_FILE_PATH)
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
