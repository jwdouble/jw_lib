package conf

import (
	"github.com/spf13/viper"

	"jw.lib/logx"
)

var vip *viper.Viper

func init() {
	vip = viper.New()

	vip.SetConfigFile(CONF_FILE_PATH)
	err := vip.ReadInConfig()
	if err != nil {
		logx.Errorf(err, "viper.ReadInConfig")
	}
}

func Get(key string) string {
	val := vip.Get(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
