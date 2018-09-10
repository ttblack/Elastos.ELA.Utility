package elalog

var Disabled *disableLog

type disableLog struct{}

func (l *disableLog) Info(e ...interface{})                  {}
func (l *disableLog) Infof(format string, e ...interface{})  {}
func (l *disableLog) Warn(e ...interface{})                  {}
func (l *disableLog) Warnf(format string, e ...interface{})  {}
func (l *disableLog) Error(e ...interface{})                 {}
func (l *disableLog) Errorf(format string, e ...interface{}) {}
func (l *disableLog) Fatal(e ...interface{})                 {}
func (l *disableLog) Fatalf(format string, e ...interface{}) {}
func (l *disableLog) Trace(e ...interface{})                 {}
func (l *disableLog) Tracef(format string, e ...interface{}) {}
func (l *disableLog) Debug(e ...interface{})                 {}
func (l *disableLog) Debugf(format string, e ...interface{}) {}
