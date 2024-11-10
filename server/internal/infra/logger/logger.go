package logger

import (
	"fmt"
	"io"
	_Log "log"
	"math/rand"
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
	logger   _Log.Logger
	Label    string
  config   *config.Config
  triggers map[int]time.Time
}

var Log Logger

func NewLogger(label string, config *config.Config) *Logger {
	log := _Log.New(os.Stdout, "", 0)
	path := "logs/" + time.Now().Format("2006-01-02-15:04:05") + ".log"

	os.MkdirAll("./logs", 0755)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{
			logger: *log,
			Label:  label,
      config: config,
      triggers: make(map[int]time.Time),
		}
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	return &Logger{
		logger: *log,
		Label:  label,
    config: config,
    triggers: make(map[int]time.Time),
	}
}

func init() {
  Log = *NewLogger("CORE", nil)
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

func (self *Logger) StartTrigger() int {
  if self.config.Debug {
    id := rand.Int()
    self.triggers[id] = time.Now()
    return id
  }

  return 0
}

func (self *Logger) Debug(args ...interface{}) {
	if self.config != nil && self.config.Debug {
    self.logWithLevel("DEBUG", Magenta, args...)
  }
}

func (self *Logger) DebugWithTime(triggerID int, args ...interface{}) {
  if self.config.Debug {
    trigger, exists := self.triggers[triggerID]
    if exists {
      args = append(args, " ELAPSED=")
      args = append(args, time.Since(trigger))
    }

    self.logWithLevel("DEBUG", Magenta, args...)
    delete(self.triggers, triggerID)
  }
}

func (self *Logger) Error(args ...interface{}) {
	self.logWithLevel("ERROR", Red, args...)
}

func (self *Logger) Warn(args ...interface{}) {
	self.logWithLevel("WARN", Yellow, args...)
}

func (self *Logger) Fatal(args ...interface{}) {
	self.logWithLevel("FATAL", RedBackground, args...)
}

func (self *Logger) Panic(args ...interface{}) {
	self.logWithLevel("PANIC", RedBackground, args...)
}

func (self *Logger) New(label string) models.Logger {
	return NewLogger(label, self.config)
}
