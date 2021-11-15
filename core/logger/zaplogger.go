package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Driver = (*ZapLogger)(nil)

// ZapLogger is a logger implementation that uses zap.
type ZapLogger struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

// NewZapLogger creates a new ZapLogger.
func NewZapLogger(opt *Option) (*ZapLogger, error) {
	filepath := fmt.Sprintf("%s/%s.log", strings.TrimRight(opt.Path, "/"), opt.Filename)

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{filepath}
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.WarnLevel),
		zap.AddCallerSkip(3),
	)
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		zap:   logger,
		sugar: logger.Sugar(),
	}, nil
}

// Write writes a message to the log.
func (l *ZapLogger) Write(level Level, args ...interface{}) error {
	switch level {
	case DebugLevel:
		l.sugar.Debug(args...)
	case InfoLevel:
		l.sugar.Info(args...)
	case WarnLevel:
		l.sugar.Warn(args...)
	case ErrorLevel:
		l.sugar.Error(args...)
	default:
		panic("logger: unknown level")
	}
	return nil
}

// Close closes the logger.
func (l *ZapLogger) Close() error {
	return l.zap.Sync()
}
