package log

import (
	"fmt"
	"time"
)

var _ Writer = &writer{}

// WriterConfig is writer's configuration
type WriterConfig struct {
	Handler Handler
	Exit    ExitFunc
}

type writer struct {
	fields  FieldMap
	exit    ExitFunc
	handler Handler
}

// NewWriter creates a new writer
func NewWriter(cfg *WriterConfig) Writer {
	return &writer{
		handler: cfg.Handler,
		exit:    cfg.Exit,
		fields:  FieldMap{},
	}
}

// Fields returns the entry
func (e *writer) Fields() FieldMap {
	return e.fields
}

// WithField returns a new log entry with the supplied field.
func (e *writer) WithField(key string, value interface{}) Writer {
	return e.WithFields(F(key, value))
}

// WithFields returns a new log entry with the supplied fields appended
func (e *writer) WithFields(entries ...Fielder) Writer {
	w := &writer{
		handler: e.handler,
		exit:    e.exit,
		fields:  e.fields.copy(),
	}

	for _, entry := range entries {
		for key, value := range entry.Fields() {
			w.fields[key] = value
		}
	}

	return w
}

// WithError add a minimal stack trace to the log writer
func (e *writer) WithError(err error) Writer {
	return e.WithFields(FieldsOfError(err))
}

// Debug logs a debug entry
func (e *writer) Debug(v ...interface{}) {
	entry := Entry{
		Message: fmt.Sprint(v...),
		Level:   DebugLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(&entry)
}

// Debugf logs a debug entry with formatting
func (e *writer) Debugf(s string, v ...interface{}) {
	entry := Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   DebugLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(&entry)
}

// Info logs a normal. information, entry
func (e *writer) Info(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   InfoLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Infof logs a normal. information, entry with formatting
func (e *writer) Infof(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   InfoLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Notice logs a notice log entry
func (e writer) Notice(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   NoticeLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Noticef logs a notice log entry with formatting
func (e writer) Noticef(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   NoticeLevel,
	}
	e.handle(entry)
}

// Warn logs a warn log entry
func (e *writer) Warn(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   WarnLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Warnf logs a warn log entry with formatting
func (e writer) Warnf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   WarnLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Panic logs a panic log entry
func (e *writer) Panic(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   PanicLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Panicf logs a panic log entry with formatting
func (e *writer) Panicf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   PanicLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Alert logs an alert log entry
func (e *writer) Alert(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   AlertLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Alertf logs an alert log entry with formatting
func (e *writer) Alertf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   AlertLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Fatal logs a fatal log entry
func (e *writer) Fatal(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   FatalLevel,
		Fields:  e.fields.copy(),
	}

	e.handle(entry)
}

// Fatalf logs a fatal log entry with formatting
func (e *writer) Fatalf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   FatalLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Error logs an error log entry
func (e *writer) Error(v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprint(v...),
		Level:   ErrorLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

// Errorf logs an error log entry with formatting
func (e *writer) Errorf(s string, v ...interface{}) {
	entry := &Entry{
		Message: fmt.Sprintf(s, v...),
		Level:   ErrorLevel,
		Fields:  e.fields.copy(),
	}
	e.handle(entry)
}

func (e *writer) handle(entry *Entry) {
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
