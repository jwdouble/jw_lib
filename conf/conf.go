package conf

import (
	"github.com/spf13/viper"

	"jw.lib/logx"
)

func Get(key string) string {
	viper.SetConfigFile(CONF_FILE_PATH)
	err := viper.ReadInConfig()
	if err != nil {
		logx.Errorf(err, "viper.ReadInConfig")
		return ""
	}

	return viper.Get(key).(string)
}
