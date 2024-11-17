package models

type Logger interface {
  Infof(string, ...any)
  Errorf(string, ...any)
  Fatalf(string, ...any)
  Warnf(string, ...any)
  Debugf(string, ...any)
  Panicf(string, ...any)
  New(string) Logger
}
