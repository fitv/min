package logger

import (
	"fmt"
	"log"
	"time"
)

type Driver interface {
	Write(level Level, args ...interface{}) error
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

func (l *Logger) Debug(args ...interface{}) {
	l.Log(DebugLevel, args)
}

func (l *Logger) Info(args ...interface{}) {
	l.Log(InfoLevel, args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Log(WarnLevel, args)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log(ErrorLevel, args)
}

// Log writes a message to the log using the given level.
func (l *Logger) Log(level Level, args ...interface{}) {
	if level < l.level {
		return
	}

	err := l.driver.Write(level, args...)
	if err != nil {
		log.Println(fmt.Errorf("write log error: %w", err))
	}
}

func (l *Logger) Close() error {
	return l.driver.Close()
}

// today returns the current date in YYYY-MM-DD format.
func today() string {
	return time.Now().Format("20060102")
}
