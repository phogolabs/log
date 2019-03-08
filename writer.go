package log

import (
	"fmt"
	"time"
)

var _ Writer = writer{}

// WriterConfig is writer's configuration
type WriterConfig struct {
	Handler Handler
	Exit    ExitFunc
}

type writer struct {
	entry   Entry
	exit    ExitFunc
	handler Handler
}

// NewWriter creates a new writer
func NewWriter(cfg *WriterConfig) Writer {
	return writer{
		handler: cfg.Handler,
		exit:    cfg.Exit,
	}
}

// Entry returns the entry
func (e writer) Entry() Entry {
	return e.entry
}

// WithField returns a new log entry with the supplied field.
func (e writer) WithField(key string, value interface{}) Writer {
	e.entry.Fields = append(e.entry.Fields, Field{Key: key, Value: value})
	return e
}

// WithFields returns a new log entry with the supplied fields appended
func (e writer) WithFields(fields ...Fielder) Writer {
	for _, kv := range fields {
		e.entry.Fields = append(e.entry.Fields, kv.Fields()...)
	}
	return e
}

// WithError add a minimal stack trace to the log writer
func (e writer) WithError(err error) Writer {
	return e.WithFields(FieldsOfError(err))
}

// Debug logs a debug entry
func (e writer) Debug(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = DebugLevel
	e.handle()
}

// Debugf logs a debug entry with formatting
func (e writer) Debugf(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = DebugLevel
	e.handle()
}

// Info logs a normal. information, entry
func (e writer) Info(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = InfoLevel
	e.handle()
}

// Infof logs a normal. information, entry with formatting
func (e writer) Infof(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = InfoLevel
	e.handle()
}

// Notice logs a notice log entry
func (e writer) Notice(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = NoticeLevel
	e.handle()
}

// Noticef logs a notice log entry with formatting
func (e writer) Noticef(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = NoticeLevel
	e.handle()
}

// Warn logs a warn log entry
func (e writer) Warn(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = WarnLevel
	e.handle()
}

// Warnf logs a warn log entry with formatting
func (e writer) Warnf(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = WarnLevel
	e.handle()
}

// Panic logs a panic log entry
func (e writer) Panic(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = PanicLevel
	e.handle()
}

// Panicf logs a panic log entry with formatting
func (e writer) Panicf(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = PanicLevel
	e.handle()
}

// Alert logs an alert log entry
func (e writer) Alert(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = AlertLevel
	e.handle()
}

// Alertf logs an alert log entry with formatting
func (e writer) Alertf(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = AlertLevel
	e.handle()
}

// Fatal logs a fatal log entry
func (e writer) Fatal(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = FatalLevel
	e.handle()
}

// Fatalf logs a fatal log entry with formatting
func (e writer) Fatalf(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = FatalLevel
	e.handle()
}

// Error logs an error log entry
func (e writer) Error(v ...interface{}) {
	e.entry.Message = fmt.Sprint(v...)
	e.entry.Level = ErrorLevel
	e.handle()
}

// Errorf logs an error log entry with formatting
func (e writer) Errorf(s string, v ...interface{}) {
	e.entry.Message = fmt.Sprintf(s, v...)
	e.entry.Level = ErrorLevel
	e.handle()
}

func (e writer) handle() {
	e.entry.Timestamp = time.Now()

	if e.handler != nil {
		e.handler.Handle(&e.entry)
	}

	if e.exit != nil {
		switch e.entry.Level {
		case PanicLevel:
			e.exit(1)
		case FatalLevel:
			e.exit(1)
		}
	}
}
