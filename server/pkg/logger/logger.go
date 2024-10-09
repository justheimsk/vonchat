package logger

import (
	"io"
	Log "log"
	"os"
	"time"
)

var log *Log.Logger

func init() {
	log = Log.Default()
	path := "logs/" + time.Now().Format("2006-01-02-15:04:05") + ".log"

	os.MkdirAll("./logs", 0755)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to create log file, logging to STDOUT only.")
		return
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.Println("Started logging to STDOUT and " + path)
}

func GetLogger() *Log.Logger {
	return log
}
