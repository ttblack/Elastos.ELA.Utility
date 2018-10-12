package elalog

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
