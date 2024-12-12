package models

type Logger interface {
	Infof(string, ...any)
	Errorf(string, ...any)
	Fatalf(string, ...any)
	Warnf(string, ...any)
	Debugf(string, ...any)
	Info(interface{}, ...interface{})
	Error(interface{}, ...interface{})
	Fatal(interface{}, ...interface{})
	Warn(interface{}, ...interface{})
	Debug(interface{}, ...interface{})
}
