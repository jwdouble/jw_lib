package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// TODO: fix
func Get(key string) string {
	viper.SetConfigFile(CONF_FILE_PATH)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return viper.Get(key).(string)
}
