package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Level represents a log level
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

// Logger is a structured logger
type Logger struct {
	level  Level
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
}

var defaultLogger *Logger

func init() {
	defaultLogger = New(os.Stderr, LevelInfo)
}

// New creates a new logger
func New(w io.Writer, level Level) *Logger {
	return &Logger{
		level: level,
		debug: log.New(w, "[DEBUG] ", log.Ltime|log.Lshortfile),
		info:  log.New(w, "[INFO]  ", log.Ltime),
		warn:  log.New(w, "[WARN]  ", log.Ltime),
		error: log.New(w, "[ERROR] ", log.Ltime|log.Lshortfile),
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level <= LevelDebug {
		l.debug.Output(2, fmt.Sprintf(format, v...))
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= LevelInfo {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level <= LevelWarn {
		l.warn.Output(2, fmt.Sprintf(format, v...))
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	if l.level <= LevelError {
		l.error.Output(2, fmt.Sprintf(format, v...))
	}
}

// Global functions using the default logger

// Debug logs a debug message using the default logger
func Debug(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

// Info logs an info message using the default logger
func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

// Warn logs a warning message using the default logger
func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

// Error logs an error message using the default logger
func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

// SetLevel sets the log level for the default logger
func SetLevel(level Level) {
	defaultLogger.level = level
}

// GetLevel returns the current log level
func GetLevel() Level {
	return defaultLogger.level
}

// ParseLevel parses a string into a log level
func ParseLevel(s string) (Level, error) {
	switch s {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	default:
		return LevelInfo, fmt.Errorf("unknown log level: %s", s)
	}
}
