package log

import (
	"os"
	"time"
)

// Writer of the log
type Writer interface {
	// WithField returns a new log entry with the supplied field.
	WithField(key string, value interface{}) Writer

	// WithFields returns a new log entry with the supplied fields appended
	WithFields(fields ...Field) Writer

	// WithFieldMap returns a new log entry with the supplied fields appended
	WithFieldMap(m map[string]interface{}) Writer

	// WithError add a minimal stack trace to the log Entry
	WithError(err error) Writer

	// Debug logs a debug entry
	Debug(v ...interface{})

	// Debugf logs a debug entry with formatting
	Debugf(s string, v ...interface{})

	// Info logs a normal. information, entry
	Info(v ...interface{})

	// Infof logs a normal. information, entry with formatting
	Infof(s string, v ...interface{})

	// Notice logs a notice log entry
	Notice(v ...interface{})

	// Noticef logs a notice log entry with formatting
	Noticef(s string, v ...interface{})

	// Warn logs a warn log entry
	Warn(v ...interface{})

	// Warnf logs a warn log entry with formatting
	Warnf(s string, v ...interface{})

	// Panic logs a panic log entry
	Panic(v ...interface{})

	// Panicf logs a panic log entry with formatting
	Panicf(s string, v ...interface{})

	// Alert logs an alert log entry
	Alert(v ...interface{})

	// Alertf logs an alert log entry with formatting
	Alertf(s string, v ...interface{})

	// Fatal logs a fatal log entry
	Fatal(v ...interface{})

	// Fatalf logs a fatal log entry with formatting
	Fatalf(s string, v ...interface{})

	// Error logs an error log entry
	Error(v ...interface{})

	// Errorf logs an error log entry with formatting
	Errorf(s string, v ...interface{})
}

// standard logger
var std = Entry{
	Timestamp: time.Now(),
	Exit:      os.Exit,
	Handler: &LevelHandler{
		Level:   DebugLevel,
		Handler: &DefaultHandler{},
	},
}

// SetExitFn sets the exit function. default: os.Exit
func SetExitFn(fn ExitFunc) {
	std.Exit = fn
}

// SetLevel sets the level handler
func SetLevel(level Level) {
	leveled := std.Handler.(*LevelHandler)
	leveled.Level = level
}

// SetHandler sets the handler
func SetHandler(handler Handler) {
	leveled := std.Handler.(*LevelHandler)
	leveled.Handler = handler
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
func WithField(key string, value interface{}) Writer {
	return std.WithField(key, value)
}

// WithFields returns a new log entry with the supplied fields appended
func WithFields(fields ...Field) Writer {
	return std.WithFields(fields...)
}

// WithFieldMap returns a new log entry with the supplied fields appended
func WithFieldMap(m map[string]interface{}) Writer {
	return std.WithFieldMap(m)
}

// WithError add a minimal stack trace to the log Entry
func WithError(err error) Writer {
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
