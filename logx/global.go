package logx

func Debugf(str string, arg ...interface{}) {
	l := getLogger()
	l.Debug().Msgf(str, arg...)
}

func Infof(str string, arg ...interface{}) {
	l := getLogger()
	l.Info().Msgf(str, arg...)
}

func Warnf(str string, arg ...interface{}) {
	l := getLogger()
	l.Warn().Msgf(str, arg...)
}

func Errorf(err error, str string, arg ...interface{}) {
	l := getLogger()
	l.Errorf(err, str, arg...)
}

func Fatalf(str string, arg ...interface{}) {
	l := getLogger()
	l.Fatal().Msgf(str, arg...)
}
