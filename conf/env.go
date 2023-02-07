package conf

import (
	"log"
	"sync"
)

type EnvVar string

const (
	COMMON_PASSWORD EnvVar = "GO_COMMON_PASSWORD"
	SERVER_PORT     EnvVar = "GO_SERVER_PORT"
)

var sm sync.Map

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Get() string {
	err := vip.BindEnv(e.String())
	if err != nil {
		return ""
	}

	return vip.GetString(e.String())
}

func (e EnvVar) Value(in string) string {
	es := e.String()

	// 预存值, 如果环境变量已配置用环境变量覆盖
	if env := e.Get(); env != "" {
		sm.Store(es, env)
	} else {
		sm.Store(es, in)
	}

	val, ok := sm.Load(es)
	if !ok {
		log.Fatalf("env var %s not found", es)
	}

	return val.(string)
}
