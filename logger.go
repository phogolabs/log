package log

import (
	"fmt"
	"time"
)

var _ Logger = &logger{}

type logger struct {
	fields  Map
	exit    ExitFunc
	handler Handler
}

// New creates a new logger
func New(cfg *Config) Logger {
	return &logger{
		handler: cfg.Handler,
		exit:    cfg.Exit,
		fields:  Map{},
	}
}

// Fields returns the entry
func (e *logger) Fields() Map {
	return e.fields
}

// WithField returns a new log entry with the supplied field.
func (e *logger) WithField(key string, value interface{}) Logger {
	return e.WithFields(Map{key: value})
}

// WithFields returns a new log entry with the supplied fields appended
func (e *logger) WithFields(fields Map) Logger {
	w := &logger{
		handler: e.handler,
		exit:    e.exit,
		fields:  e.copy(),
	}

	for key, value := range fields {
		w.fields[key] = value
	}

	return w
}

// WithError add a minimal stack trace to the log logger
func (e *logger) WithError(err error) Logger {
	return e.WithFields(FieldsOfError(err))
}

// Debug logs a debug entry
func (e *logger) Debug(v ...interface{}) {
	e.handle(e.entryv(DebugLevel, v))
}

// Debugf logs a debug entry with formatting
func (e *logger) Debugf(s string, v ...interface{}) {
	e.handle(e.entryf(DebugLevel, s, v))
}

// Info logs a normal. information, entry
func (e *logger) Info(v ...interface{}) {
	e.handle(e.entryv(InfoLevel, v))
}

// Infof logs a normal. information, entry with formatting
func (e *logger) Infof(s string, v ...interface{}) {
	e.handle(e.entryf(InfoLevel, s, v))
}

// Notice logs a notice log entry
func (e logger) Notice(v ...interface{}) {
	e.handle(e.entryv(NoticeLevel, v))
}

// Noticef logs a notice log entry with formatting
func (e logger) Noticef(s string, v ...interface{}) {
	e.handle(e.entryf(NoticeLevel, s, v))
}

// Warn logs a warn log entry
func (e *logger) Warn(v ...interface{}) {
	e.handle(e.entryv(WarnLevel, v))
}

// Warnf logs a warn log entry with formatting
func (e logger) Warnf(s string, v ...interface{}) {
	e.handle(e.entryf(WarnLevel, s, v))
}

// Panic logs a panic log entry
func (e *logger) Panic(v ...interface{}) {
	e.handle(e.entryv(PanicLevel, v))
}

// Panicf logs a panic log entry with formatting
func (e *logger) Panicf(s string, v ...interface{}) {
	e.handle(e.entryf(PanicLevel, s, v))
}

// Alert logs an alert log entry
func (e *logger) Alert(v ...interface{}) {
	e.handle(e.entryv(AlertLevel, v))
}

// Alertf logs an alert log entry with formatting
func (e *logger) Alertf(s string, v ...interface{}) {
	e.handle(e.entryf(AlertLevel, s, v))
}

// Fatal logs a fatal log entry
func (e *logger) Fatal(v ...interface{}) {
	e.handle(e.entryv(FatalLevel, v))
}

// Fatalf logs a fatal log entry with formatting
func (e *logger) Fatalf(s string, v ...interface{}) {
	e.handle(e.entryf(FatalLevel, s, v))
}

// Error logs an error log entry
func (e *logger) Error(v ...interface{}) {
	e.handle(e.entryv(ErrorLevel, v))
}

// Errorf logs an error log entry with formatting
func (e *logger) Errorf(s string, v ...interface{}) {
	e.handle(e.entryf(ErrorLevel, s, v))
}

func (e *logger) handle(entry *Entry) {
	if e.handler != nil {
		e.handler.Handle(entry)
	}

	if e.exit == nil {
		return
	}

	switch entry.Level {
	case PanicLevel:
		e.exit(1)
	case FatalLevel:
		e.exit(1)
	}
}

func (e *logger) entryf(level Level, msg string, v []interface{}) *Entry {
	entry := e.entry(level)
	entry.Message = fmt.Sprintf(msg, v...)
	return entry
}

func (e *logger) entryv(level Level, v []interface{}) *Entry {
	entry := e.entry(level)
	entry.Message = fmt.Sprint(v...)
	return entry
}

func (e *logger) entry(level Level) *Entry {
	entry := &Entry{
		Timestamp: time.Now(),
		Level:     level,
		Fields:    e.copy(),
	}

	return entry
}

func (e *logger) copy() Map {
	fields := Map{}

	for key, value := range e.fields {
		fields[key] = value
	}

	return fields
}
