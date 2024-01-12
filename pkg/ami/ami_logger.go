package ami

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents different log levels.
type LogLevel string

// Log levels
const (
	Info  LogLevel = "INFO"
	Debug LogLevel = "DEBUG"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
	Fatal LogLevel = "FATAL"
)

// Logger represents a JSON logger.
type Logger struct {
	Level     LogLevel `json:"level"`
	IsEnabled bool     `json:"enabled"`
}

// LoggerEntry represents a log entry.
type LoggerEntry struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Level     LogLevel  `json:"level,omitempty"`
	Message   string    `json:"message,omitempty"`
}

// NewLogger creates a new logger with the specified log level.
func NewLogger(level LogLevel) *Logger {
	return &Logger{Level: level, IsEnabled: true}
}

var i *Logger = nil
var w *Logger = nil
var e *Logger = nil
var d *Logger = nil
var f *Logger = nil

func D() *Logger {
	if d != nil {
		return d
	}
	d = NewLogger(Debug)
	return d
}

func I() *Logger {
	if i != nil {
		return i
	}
	i = NewLogger(Info)
	return i
}

func W() *Logger {
	if w != nil {
		return w
	}
	w = NewLogger(Warn)
	return w
}

func E() *Logger {
	if e != nil {
		return e
	}
	e = NewLogger(Error)
	return e
}

func F() *Logger {
	if f != nil {
		return f
	}
	f = NewLogger(Fatal)
	return f
}

func (l *Logger) JsonString(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Printf(err.Error())
		return ""
	}
	return string(result)
}

// log logs a message with the given level.
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if !l.IsEnabled {
		return
	}
	if l.Level <= level {
		entry := LoggerEntry{
			Timestamp: time.Now(),
			Level:     level,
			Message:   fmt.Sprintf(format, args...),
		}
		v, err := json.Marshal(entry)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling log entry: %v\n", err)
			return
		}
		fmt.Println(string(v))
		if level == Fatal {
			os.Exit(1) // or using log.Fatalf
		}
	}
}

// Info logs a message with info level.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(Info, format, args...)
}

// Debug logs a message with debug level.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(Debug, format, args...)
}

// Warn logs a message with warn level.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(Warn, format, args...)
}

// Error logs a message with error level.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(Error, format, args...)
}

// Fatal logs a message with fatal level and exits the application.
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(Fatal, format, args...)
}

// Info logs a message with info level.
func (l *Logger) InfoR(format string, args ...interface{}) *Logger {
	l.log(Info, format, args...)
	return l
}

// Debug logs a message with debug level.
func (l *Logger) DebugR(format string, args ...interface{}) *Logger {
	l.log(Debug, format, args...)
	return l
}

// Warn logs a message with warn level.
func (l *Logger) WarnR(format string, args ...interface{}) *Logger {
	l.log(Warn, format, args...)
	return l
}

// Error logs a message with error level.
func (l *Logger) ErrorR(format string, args ...interface{}) *Logger {
	l.log(Error, format, args...)
	return l
}

// Fatal logs a message with fatal level and exits the application.
func (l *Logger) FatalR(format string, args ...interface{}) *Logger {
	l.log(Fatal, format, args...)
	return l
}

func (l *Logger) SetEnabled(value bool) *Logger {
	l.IsEnabled = value
	return l
}
