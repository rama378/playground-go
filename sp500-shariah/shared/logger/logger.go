package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Level int

const (
	InfoLevel Level = iota
	WarmLevel
	ErrorLevel
)

var currentLevel = InfoLevel
var stdLogger = log.New(os.Stdout, "", 0)

func SetLevel(level Level) {
	currentLevel = level
}

func logf(level string, msg string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	out := fmt.Sprintf("[%s] [%s] %s", timestamp, level, fmt.Sprintf(msg, args...))
	stdLogger.Println(out)
}

func Info(msg string, args ...interface{}) {
	if currentLevel <= InfoLevel {
		logf("INFO", msg, args...)
	}
}

func Warm(msg string, args ...interface{}) {
	if currentLevel <= WarmLevel {
		logf("WARM", msg, args...)
	}
}

func Error(msg string, args ...interface{}) {
	if currentLevel <= ErrorLevel {
		logf("ERROR", msg, args...)
	}
}
