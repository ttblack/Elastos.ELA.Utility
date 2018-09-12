package elalog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	White  = "1;00"
	Blue   = "1;34"
	Red    = "1;31"
	Green  = "1;32"
	Yellow = "1;33"
	Cyan   = "1;36"
	Pink   = "1;35"
)

func color(code, msg string) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, msg)
}

var (
	InfoPrefix  = color(White, "[INFO]")
	WarnPrefix  = color(Yellow, "[WARN]")
	ErrorPrefix = color(Red, "[ERROR]")
	FatalPrefix = color(Pink, "[FATAL]")
	TracePrefix = color(Cyan, "[TRACE]")
	DebugPrefix = color(Green, "[DEBUG]")
)

var Stdout *stdout

type stdout struct {
	logger *log.Logger
}

func gid() uint64 {
	var buf [64]byte
	b := buf[:runtime.Stack(buf[:], false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func (s *stdout) print(prefix string, a ...interface{}) error {
	a = append([]interface{}{fmt.Sprintf("%s %s %d,", prefix, "GID", gid())}, a...)
	return s.logger.Output(2, fmt.Sprintln(a...))
}

func (s *stdout) printf(prefix, format string, a ...interface{}) error {
	a = append([]interface{}{prefix, "GID", gid()}, a...)
	return s.logger.Output(2, fmt.Sprintf("%s %s %d, "+format+"\n", a...))
}

func (s *stdout) Info(e ...interface{}) {
	s.print(InfoPrefix, e...)
}

func (s *stdout) Infof(format string, e ...interface{}) {
	s.printf(InfoPrefix, format, e...)
}

func (s *stdout) Warn(e ...interface{}) {
	s.print(WarnPrefix, e...)
}

func (s *stdout) Warnf(format string, e ...interface{}) {
	s.printf(WarnPrefix, format, e...)
}

func (s *stdout) Error(e ...interface{}) {
	s.print(ErrorPrefix, e...)
}

func (s *stdout) Errorf(format string, e ...interface{}) {
	s.printf(ErrorPrefix, format, e...)
}

func (s *stdout) Fatal(e ...interface{}) {
	s.print(FatalPrefix, e...)
}

func (s *stdout) Fatalf(format string, e ...interface{}) {
	s.printf(FatalPrefix, format, e...)
}

func (s *stdout) Trace(e ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	nameFull := f.Name()
	nameEnd := filepath.Ext(nameFull)
	funcName := strings.TrimPrefix(nameEnd, ".")

	e = append([]interface{}{funcName + "()", fileName + ":" + strconv.Itoa(line)}, e...)

	s.print(TracePrefix, e...)
}

func (s *stdout) Tracef(format string, e ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	nameFull := f.Name()
	nameEnd := filepath.Ext(nameFull)
	funcName := strings.TrimPrefix(nameEnd, ".")

	e = append([]interface{}{funcName, fileName, line}, e...)

	s.printf(TracePrefix, "%s() %s:%d "+format, e...)
}

func (s *stdout) Debug(e ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	e = append([]interface{}{f.Name(), fileName + ":" + strconv.Itoa(line)}, e...)

	s.print(DebugPrefix, e...)
}

func (s *stdout) Debugf(format string, e ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	e = append([]interface{}{f.Name(), fileName, line}, e...)

	s.printf(DebugPrefix, "%s %s:%d "+format, e...)
}

func newStdout() *stdout {
	return &stdout{
		logger: log.New(os.Stdout, "*", log.LstdFlags|log.Lmicroseconds),
	}
}
