package log

import (
	"os"
	"time"
)

// standard logger
var std = Entry{
	Timestamp: time.Now(),
	Level:     InfoLevel,
	Exit:      os.Exit,
}

// SetExitFn sets the exit function. default: os.Exit
func SetExitFn(fn ExitFunc) {
	std.Exit = fn
}

// SetHandler sets the handler
func SetHandler(handler Handler) {
	std.Handler = handler
}

// SetDefaultFields sets the default fields
func SetDefaultFields(field ...Field) {
	std.Fields = field
}

// SetDefaultFieldsWithMap sets the default fields
func SetDefaultFieldsWithMap(m map[string]interface{}) {
	std.Fields = M(m).Fields()
}

// WithField returns a new log entry with the supplied field.
func WithField(key string, value interface{}) Entry {
	return std.WithField(key, value)
}

// WithFields returns a new log entry with the supplied fields appended
func WithFields(fields ...Field) Entry {
	return std.WithFields(fields...)
}

// WithFieldMap returns a new log entry with the supplied fields appended
func WithFieldMap(m map[string]interface{}) Entry {
	return std.WithFieldMap(m)
}

// WithError add a minimal stack trace to the log Entry
func WithError(err error) Entry {
	return std.WithError(err)
}

// Debug logs a debug entry
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Debugf logs a debug entry with formatting
func Debugf(s string, v ...interface{}) {
	std.Debugf(s, v...)
}

// Info logs a normal. information, entry
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof logs a normal. information, entry with formatting
func Infof(s string, v ...interface{}) {
	std.Infof(s, v...)
}

// Notice logs a notice log entry
func Notice(v ...interface{}) {
	std.Notice(v...)
}

// Noticef logs a notice log entry with formatting
func Noticef(s string, v ...interface{}) {
	std.Noticef(s, v...)
}

// Warn logs a warn log entry
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Warnf logs a warn log entry with formatting
func Warnf(s string, v ...interface{}) {
	std.Warnf(s, v...)
}

// Panic logs a panic log entry
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// Panicf logs a panic log entry with formatting
func Panicf(s string, v ...interface{}) {
	std.Panicf(s, v...)
}

// Alert logs an alert log entry
func Alert(v ...interface{}) {
	std.Alert(v...)
}

// Alertf logs an alert log entry with formatting
func Alertf(s string, v ...interface{}) {
	std.Alertf(s, v...)
}

// Fatal logs a fatal log entry
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf logs a fatal log entry with formatting
func Fatalf(s string, v ...interface{}) {
	std.Fatalf(s, v...)
}

// Error logs an error log entry
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf logs an error log entry with formatting
func Errorf(s string, v ...interface{}) {
	std.Errorf(s, v...)
}
