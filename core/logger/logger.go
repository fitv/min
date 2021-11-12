package logger

import (
	"fmt"
	"log"
	"time"
)

type Driver interface {
	Write([]byte) (int, error)
	Close() error
}

type Logger struct {
	level  Level
	driver Driver
}

type Option struct {
	Filename string
	Path     string
	Daily    bool
}

func New(level Level, driver Driver) *Logger {
	return &Logger{
		level:  level,
		driver: driver,
	}
}

func (l *Logger) Debug(a ...interface{}) {
	l.Log(DebugLevel, a)
}

func (l *Logger) Info(a ...interface{}) {
	l.Log(InfoLevel, a)
}

func (l *Logger) Warn(a ...interface{}) {
	l.Log(WarnLevel, a)
}

func (l *Logger) Error(a ...interface{}) {
	l.Log(ErrorLevel, a)
}

// Log writes a message to the log using the given level.
func (l *Logger) Log(level Level, a ...interface{}) {
	if level < l.level {
		return
	}

	datetime := time.Now().Format("2006/01/02 15:04:05")
	a = append([]interface{}{datetime, fmt.Sprintf("[%s]", level)}, a...)

	_, err := l.driver.Write([]byte(fmt.Sprintln(a...)))
	if err != nil {
		log.Println(fmt.Errorf("write log error: %w", err))
	}
}

func (l *Logger) Close() error {
	return l.driver.Close()
}
