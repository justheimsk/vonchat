package logger

import (
	"fmt"
	"io"
	Log "log"
	"os"
	"time"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

const (
	Reset         = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	White         = "\033[37m"
	RedBackground = "\033[41m"
)

type Logger struct {
	logger Log.Logger
	Label  string
}

func NewLogger(label string) *Logger {
	log := Log.New(os.Stdout, "", 0)
	path := "logs/" + time.Now().Format("2006-01-02-15:04:05") + ".log"

	os.MkdirAll("./logs", 0755)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{
			logger: *log,
			Label:  label,
		}
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	return &Logger{
		logger: *log,
		Label:  label,
	}
}

func (self *Logger) logWithLevel(level string, color string, args ...interface{}) {
	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Fprintf(self.logger.Writer(), "%s [ %s ]"+color+" %s "+Reset, now, self.Label, level)

	switch level {
	case "FATAL":
		self.logger.Fatal(args...)
	case "PANIC":
		self.logger.Panic(args...)
	default:
		self.logger.Print(args...)
	}
}

func (self *Logger) Info(args ...interface{}) {
	self.logWithLevel("INFO", Blue, args...)
}

func (self *Logger) Error(args ...interface{}) {
	self.logWithLevel("ERROR", Red, args...)
}

func (self *Logger) Fatal(args ...interface{}) {
	self.logWithLevel("FATAL", RedBackground, args...)
}

func (self *Logger) Panic(args ...interface{}) {
	self.logWithLevel("PANIC", RedBackground, args...)
}

func (self *Logger) New(label string) models.Logger {
	return NewLogger(label)
}
