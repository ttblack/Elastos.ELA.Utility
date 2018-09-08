package elalog

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	*zap.SugaredLogger
}

func NewZapLogger() Logger {
	logger, _ := zap.NewProduction()
	return &ZapLogger{
		SugaredLogger:logger.Sugar(),
	}
}

func (z *ZapLogger) Info(e ...interface{}) {
	defer z.Sync()
	z.Info(e)
}

func (z *ZapLogger) Infof(format string, e ...interface{}) {
	defer z.Sync()
	z.Infof(format, e)
}

func (z *ZapLogger) Warn(e ...interface{}) {
	defer z.Sync()
	z.Warn(e)
}

func (z *ZapLogger) Warnf(format string, e ...interface{}) {
	defer z.Sync()
	z.Warnf(format, e)
}

func (z *ZapLogger) Error(e ...interface{}) {
	defer z.Sync()
	z.Error(e)
}

func (z *ZapLogger) Errorf(format string, e ...interface{}) {
	defer z.Sync()
	z.Errorf(format, e)
}


func (z *ZapLogger) Fatal(e ...interface{}) {
	defer z.Sync()
	z.Fatal(e)
}

func (z *ZapLogger) Fatalf(format string, e ...interface{}) {
	defer z.Sync()
	z.Fatalf(format, e)
}


func (z *ZapLogger) Trace(e ...interface{}) {
	defer z.Sync()
	z.Trace(e)
}

func (z *ZapLogger) Tracef(format string, e ...interface{}) {
	defer z.Sync()
	z.Tracef(format, e)
}


func (z *ZapLogger) Debug(e ...interface{}) {
	defer z.Sync()
	z.Debug(e)
}

func (z *ZapLogger) Debugf(format string, e ...interface{}) {
	defer z.Sync()
	z.Debugf(format, e)
}
