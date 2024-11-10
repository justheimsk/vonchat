package models

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
  Debug(...interface{})
  DebugWithTime(int, ...interface{})
  StartTrigger() int
	Panic(...interface{})
	New(string) Logger
}
