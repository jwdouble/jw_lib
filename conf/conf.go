package conf

import (
	"github.com/spf13/viper"
)

var vip *viper.Viper

func init() {
	vip = viper.New()

	vip.SetConfigFile(CONF_FILE_PATH)
	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	val := vip.Get(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
