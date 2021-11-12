package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var _ Driver = (*FileLogger)(nil)

type FileLogger struct {
	mux      sync.Mutex
	file     *os.File
	filename string
	path     string
	date     string
	daily    bool
}

// NewFileLogger creates a new FileLogger.
func NewFileLogger(opt *Option) *FileLogger {
	logger := &FileLogger{
		path:     strings.TrimRight(opt.Path, "/"),
		filename: opt.Filename,
		daily:    opt.Daily,
	}
	if logger.daily {
		logger.date = today()
	}
	return logger
}

// WithFields adds fields to the logger.
func (l *FileLogger) Write(p []byte) (n int, err error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if l.file == nil {
		if err := l.openFile(); err != nil {
			return 0, err
		}
	}
	if l.daily && today() != l.date {
		if err := l.close(); err != nil {
			return 0, err
		}
		l.date = today()
		if err := l.openFile(); err != nil {
			l.date = ""
			return 0, err
		}
	}

	return l.file.Write(p)
}

// Close closes the logger.
func (l *FileLogger) Close() error {
	l.mux.Lock()
	defer l.mux.Unlock()
	return l.close()
}

func (l *FileLogger) close() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil
	return err
}

// openFile opens the log file.
func (l *FileLogger) openFile() error {
	filename := fmt.Sprintf("%s/%s.log", l.path, l.filename)
	if l.daily {
		filename = fmt.Sprintf("%s/%s-%s.log", l.path, l.filename, l.date)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	l.file = file
	return nil
}

// today returns the current date in YYYY-MM-DD format.
func today() string {
	return time.Now().Format("20060102")
}
