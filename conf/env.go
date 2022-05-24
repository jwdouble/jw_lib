package conf

import (
	"github.com/spf13/viper"
	"sync"
)

type EnvVar string

const CONF_FILE_PATH = "../config.yaml"

const (
	APP_PORT EnvVar = "APP_PORT"
)

var sm sync.Map

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Value(v string) string {
	es := e.String()
	// 预存值, 如果环境变量已配置用环境变量覆盖
	v1 := e.GetEnv()
	if v1 != "" {
		sm.Store(es, v1)
	} else {
		sm.Store(es, v)
	}

	val, ok := sm.Load(es)
	if !ok {
		panic("env key: " + es + " not exist")
	}
	return val.(string)
}

func (e EnvVar) GetEnv() string {
	viper.SetEnvPrefix("k8s")
	err := viper.BindEnv(e.String())
	if err != nil {
		panic(err)
	}

	return viper.GetString(e.String())
}
