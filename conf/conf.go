package conf

import "github.com/spf13/viper"

func Get(key string) string {
	viper.SetConfigFile(CONF_FILE_PATH)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.Get(key).(string)
}
