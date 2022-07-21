package logx

func Errorf(err error, str string, arg ...interface{}) {
	l := getLogger()
	l.Errorf(err, str, arg...)
}
