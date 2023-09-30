package logger

import "github.com/dehwyy/logd"

type AppLogger interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

func New() AppLogger {
	return logd.New()
}
