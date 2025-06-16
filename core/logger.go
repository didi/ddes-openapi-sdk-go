package core

import (
	"context"
	"log"
	"os"
	"strings"
)

type LogLevel int // 日志等级

const (
	LogLevelDebug LogLevel = iota + 1
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
}

type LoggerImpl struct {
	LogLevel LogLevel
	Logger   Logger
}

func NewLoggerImpl(level LogLevel, logger Logger) *LoggerImpl {
	return &LoggerImpl{
		LogLevel: level,
		Logger:   logger,
	}
}

func (logger *LoggerImpl) Debug(ctx context.Context, args ...interface{}) {
	if nil != logger.Logger && logger.LogLevel <= LogLevelDebug {
		logger.Logger.Debug(ctx, args...)
	}
}

func (logger *LoggerImpl) Info(ctx context.Context, args ...interface{}) {
	if nil != logger.Logger && logger.LogLevel <= LogLevelInfo {
		logger.Logger.Info(ctx, args...)
	}
}

func (logger *LoggerImpl) Warn(ctx context.Context, args ...interface{}) {
	if nil != logger.Logger && logger.LogLevel <= LogLevelWarn {
		logger.Logger.Warn(ctx, args...)
	}
}

func (logger *LoggerImpl) Error(ctx context.Context, args ...interface{}) {
	if nil != logger.Logger && logger.LogLevel <= LogLevelError {
		logger.Logger.Error(ctx, args...)
	}
}

func NewDefaultLogger(level LogLevel) Logger {
	return &LoggerImpl{
		LogLevel: level,
		Logger:   &defaultLogger{logger: log.New(os.Stdout, "", log.Ldate|log.Ltime)},
	}
}

type defaultLogger struct {
	logger *log.Logger
}

func (l *defaultLogger) Debug(ctx context.Context, args ...interface{}) {
	l.logger.SetPrefix("[Debug] ")
	format := make([]string, len(args))
	for i := range format {
		format[i] = "%v"
	}
	l.logger.Printf(strings.Join(format, " "), args...)
}

func (l *defaultLogger) Info(ctx context.Context, args ...interface{}) {
	l.logger.SetPrefix("[Info] ")
	format := make([]string, len(args))
	for i := range format {
		format[i] = "%v"
	}
	l.logger.Printf(strings.Join(format, " "), args...)
}

func (l *defaultLogger) Warn(ctx context.Context, args ...interface{}) {
	l.logger.SetPrefix("[Warn] ")
	format := make([]string, len(args))
	for i := range format {
		format[i] = "%v"
	}
	l.logger.Printf(strings.Join(format, " "), args...)
}

func (l *defaultLogger) Error(ctx context.Context, args ...interface{}) {
	l.logger.SetPrefix("[Error] ")
	format := make([]string, len(args))
	for i := range format {
		format[i] = "%v"
	}
	l.logger.Printf(strings.Join(format, " "), args...)
}
