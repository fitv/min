package logger

import (
	"fmt"
	"log"
	"time"
)

// Driver is the logger interface.
type Driver interface {
	Write(level Level, args ...interface{}) error
	Close() error
}

// Logger is the logger wrapper.
type Logger struct {
	level  Level
	driver Driver
}

// OPtion is the option for logger.
type Option struct {
	Filename string
	Path     string
	Daily    bool
}

// New creates a new Logger.
func New(level Level, driver Driver) *Logger {
	return &Logger{
		level:  level,
		driver: driver,
	}
}

// Debug writes a message to the log using the DEBUG level.
func (l *Logger) Debug(args ...interface{}) {
	l.Log(DebugLevel, args)
}

// Info writes a message to the log using the INFO level.
func (l *Logger) Info(args ...interface{}) {
	l.Log(InfoLevel, args)
}

// Warn writes a message to the log using the WARN level.
func (l *Logger) Warn(args ...interface{}) {
	l.Log(WarnLevel, args)
}

// Error writes a message to the log using the ERROR level.
func (l *Logger) Error(args ...interface{}) {
	l.Log(ErrorLevel, args)
}

// Log writes a message to the log using the given level.
func (l *Logger) Log(level Level, args ...interface{}) {
	// Skip if the level is below the logger level.
	if level < l.level {
		return
	}

	err := l.driver.Write(level, args...)
	// when write log to file failed, print to console
	if err != nil {
		log.Println(fmt.Errorf("write log error: %w", err))
	}
}

// Close closes the logger.
func (l *Logger) Close() error {
	return l.driver.Close()
}

// today returns the current date in YYYYMMDD format.
func today() string {
	return time.Now().Format("20060102")
}
