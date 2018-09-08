package log

var logger Logger

type Logger interface {
	Info(e ...interface{})
	Infof(format string, e ...interface{})

	Warn(e ...interface{})
	Warnf(format string, e ...interface{})

	Error(e ...interface{})
	Errorf(format string, e ...interface{})

	Fatal(e ...interface{})
	Fatalf(format string, e ...interface{})

	Trace(e ...interface{})
	Tracef(format string, e ...interface{})

	Debug(e ...interface{})
	Debugf(format string, e ...interface{})
}

func Info(e ...interface{}) {
	logger.Info(e)
}

func Infof(format string, e ...interface{}) {
	logger.Infof(format, e)
}

func Warn(e ...interface{}) {
	logger.Warn(e)
}

func Warnf(format string, e ...interface{}) {
	logger.Warnf(format, e)
}

func Error(e ...interface{}) {
	logger.Error(e)
}

func Errorf(format string, e ...interface{}) {
	logger.Errorf(format, e)
}

func Fatal(e ...interface{}) {
	logger.Fatal(e)
}

func Fatalf(format string, e ...interface{}) {
	logger.Fatalf(format, e)
}

func Trace(e ...interface{}) {
	logger.Trace(e)
}

func Tracef(format string, e ...interface{}) {
	logger.Tracef(format, e)
}

func Debug(e ...interface{}) {
	logger.Debug(e)
}

func Debugf(format string, e ...interface{}) {
	logger.Debugf(format, e)
}

func SetLogger(log Logger) {
	logger = log
}