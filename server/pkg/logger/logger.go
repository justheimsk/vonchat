package logger

import (
	"io"
	_log "log"
	"os"
	"time"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
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
	log    *_log.Logger
	Label  string
	config *config.Config
	file   *os.File
}

func NewLogger(label string, config *config.Config, file *os.File) models.Logger {
	log := _log.New(os.Stdout, "", 0)
	path := "logs/" + time.Now().Format("2006-01-02-15:04:05") + ".log"

	logFile := file
	if logFile == nil {
		_ = os.MkdirAll("./logs", 0755)
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return &Logger{
				log:    log,
				Label:  label,
				config: config,
			}
		}

		logFile = file
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	return &Logger{
		log:    log,
		Label:  label,
		config: config,
		file:   logFile,
	}
}

func (self *Logger) concatLevel(format string, level string, color string) string {
	now := time.Now().Format("2006-01-02-15:04:05")
	return now + " [ " + self.Label + " ] " + color + level + Reset + " " + format
}

func (self *Logger) Infof(format string, args ...any) {
	self.log.Printf(self.concatLevel(format, "INFO", Blue), args...)
}

func (self *Logger) Warnf(format string, args ...any) {
	self.log.Printf(self.concatLevel(format, "WARN", Yellow), args...)
}

func (self *Logger) Debugf(format string, args ...any) {
	if self.config != nil && self.config.Debug {
		self.log.Printf(self.concatLevel(format, "DEBUG", Magenta), args...)
	}
}

func (self *Logger) Errorf(format string, args ...any) {
	self.log.Printf(self.concatLevel(format, "ERROR", Red), args...)
}

func (self *Logger) Panicf(format string, args ...any) {
	self.log.Panicf(self.concatLevel(format, "PANIC", RedBackground), args...)
}

func (self *Logger) Fatalf(format string, args ...any) {
	self.log.Fatalf(self.concatLevel(format, "FATAL", RedBackground), args...)
}

func (self *Logger) New(label string) models.Logger {
	return NewLogger(label, self.config, self.file)
}
