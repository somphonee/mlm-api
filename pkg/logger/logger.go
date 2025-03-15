package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Logger levels
const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
)

var (
	// Default logger instance
	defaultLogger *Logger
	// Output destination
	logOutput io.Writer = os.Stdout
	// Include file and line number information in logs
	includeFileInfo = true
)

// Logger represents a custom logger
type Logger struct {
	prefix string
}

// init initializes the default logger
func init() {
	defaultLogger = NewLogger("")
}

// NewLogger creates a new logger with the given prefix
func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
	}
}

// SetOutput sets the output destination for all loggers
func SetOutput(w io.Writer) {
	logOutput = w
	log.SetOutput(w)
}

// SetLogFile sets a file as the output destination
func SetLogFile(filePath string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Open log file
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Set output to both file and stdout
	mw := io.MultiWriter(os.Stdout, f)
	SetOutput(mw)
	return nil
}

// SetIncludeFileInfo sets whether to include file and line information in logs
func SetIncludeFileInfo(include bool) {
	includeFileInfo = include
}

// formatLog formats the log message with timestamp, level, and caller info
func (l *Logger) formatLog(level, msg string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	prefix := ""
	if l.prefix != "" {
		prefix = "[" + l.prefix + "] "
	}

	fileInfo := ""
	if includeFileInfo {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			file = filepath.Base(file)
			fileInfo = fmt.Sprintf("[%s:%d] ", file, line)
		}
	}

	return fmt.Sprintf("%s [%s] %s%s%s", timestamp, level, prefix, fileInfo, msg)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(logOutput, l.formatLog(LevelDebug, msg))
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(logOutput, l.formatLog(LevelInfo, msg))
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(logOutput, l.formatLog(LevelWarn, msg))
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(logOutput, l.formatLog(LevelError, msg))
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(logOutput, l.formatLog(LevelFatal, msg))
	os.Exit(1)
}

// Default logger methods
func Debug(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	defaultLogger.Fatal(format, v...)
}