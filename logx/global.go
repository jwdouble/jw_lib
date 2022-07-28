package logx

func Debugf(str string, arg ...interface{}) {
	l := getLogger()
	l.Debugf(str, arg...)
}

func Infof(str string, arg ...interface{}) {
	l := getLogger()
	l.Infof(str, arg...)
}

func Warnf(str string, arg ...interface{}) {
	l := getLogger()
	l.Warnf(str, arg...)
}

func Errorf(err error, str string, arg ...interface{}) {
	l := getErrLogger()
	l.Errorf(err, str, arg...)
}

func Fatalf(str string, arg ...interface{}) {
	l := getLogger()
	l.Fatalf(str, arg...)
}

func KV(key string, value interface{}) {
	l := getLogger()
	l.KV(key, value)
}
