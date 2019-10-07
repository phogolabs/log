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
		fields:  e.fields.copy(),
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
	entry := Entry{
		Message: fmt.Sprint(v...),
		Level:   DebugLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(&entry)
}

// Debugf logs a debug entry with formatting
func (e *logger) Debugf(s string, v ...interface{}) {
	entry := Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   DebugLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(&entry)
}

// Info logs a normal. information, entry
func (e *logger) Info(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   InfoLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Infof logs a normal. information, entry with formatting
func (e *logger) Infof(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   InfoLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Notice logs a notice log entry
func (e logger) Notice(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   NoticeLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Noticef logs a notice log entry with formatting
func (e logger) Noticef(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   NoticeLevel,
	}
	e.handle(entry)
}

// Warn logs a warn log entry
func (e *logger) Warn(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   WarnLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Warnf logs a warn log entry with formatting
func (e logger) Warnf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   WarnLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Panic logs a panic log entry
func (e *logger) Panic(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   PanicLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Panicf logs a panic log entry with formatting
func (e *logger) Panicf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   PanicLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Alert logs an alert log entry
func (e *logger) Alert(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   AlertLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Alertf logs an alert log entry with formatting
func (e *logger) Alertf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   AlertLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Fatal logs a fatal log entry
func (e *logger) Fatal(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   FatalLevel,
		Fields:  e.fields.copy(),
	}

	e.handle(entry)
}

// Fatalf logs a fatal log entry with formatting
func (e *logger) Fatalf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   FatalLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Error logs an error log entry
func (e *logger) Error(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   ErrorLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Errorf logs an error log entry with formatting
func (e *logger) Errorf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   ErrorLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

func (e *logger) handle(entry *Entry) {
	entry.Timestamp = time.Now()

	if e.handler != nil {
		e.handler.Handle(entry)
	}

	if e.exit != nil {
		switch entry.Level {
		case PanicLevel:
			e.exit(1)
		case FatalLevel:
			e.exit(1)
		}
	}
}
