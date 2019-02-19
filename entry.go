package log

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// ExitFunc is a function called on Panic or Fatal level
type ExitFunc func(code int)

// Field is a single Field key and value
type Field struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// F creates a new Field using the supplied key + value.
// it is shorthand for defining field manually
func F(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// M is a map
type M map[string]interface{}

// Fields return the map as slice of fields
func (m M) Fields() []Field {
	fields := []Field{}

	for key, value := range m {
		fields = append(fields, F(key, value))
	}

	return fields
}

//go:generate counterfeiter -fake-name Handler -o ./fake/handler.go . Handler

// Handler handles an entry
type Handler interface {
	Handle(e *Entry)
}

var _ Handler = CompositeHandler{}

// CompositeHandler is a slice of handler
type CompositeHandler []Handler

// Handle handles the entry
func (handlers CompositeHandler) Handle(e *Entry) {
	for _, handler := range handlers {
		handler.Handle(e)
	}
}

var _ Handler = &LevelHandler{}

// LevelHandler handles entries for given level
type LevelHandler struct {
	Level   Level
	Handler Handler
}

// Handle handles the entry
func (h *LevelHandler) Handle(e *Entry) {
	if e.Level < h.Level {
		return
	}

	h.Handler.Handle(e)
}

var _ Handler = &DefaultHandler{}

// DefaultHandler represents the default handler
type DefaultHandler struct{}

// Handle handles the entry
func (h *DefaultHandler) Handle(e *Entry) {
	data, _ := json.Marshal(e)
	log.SetFlags(0)
	log.Println(string(data))
}

var _ Writer = Entry{}

// Entry defines a single log entry
type Entry struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Fields    []Field   `json:"fields,omitempty"`
	Level     Level     `json:"level"`
	Exit      ExitFunc  `json:"-"`
	Handler   Handler   `json:"-"`
}

// WithField returns a new log entry with the supplied field.
func (e Entry) WithField(key string, value interface{}) Writer {
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

// WithFields returns a new log entry with the supplied fields appended
func (e Entry) WithFields(fields ...Field) Writer {
	e.Fields = append(e.Fields, fields...)
	return e
}

// WithFieldMap returns a new log entry with the supplied fields appended
func (e Entry) WithFieldMap(m map[string]interface{}) Writer {
	e.Fields = append(e.Fields, M(m).Fields()...)
	return e
}

// WithError add a minimal stack trace to the log Entry
func (e Entry) WithError(err error) Writer {
	return e.WithFields(FieldsOfError(err)...)
}

// Debug logs a debug entry
func (e Entry) Debug(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = DebugLevel
	e.handle()
}

// Debugf logs a debug entry with formatting
func (e Entry) Debugf(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = DebugLevel
	e.handle()
}

// Info logs a normal. information, entry
func (e Entry) Info(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = InfoLevel
	e.handle()
}

// Infof logs a normal. information, entry with formatting
func (e Entry) Infof(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = InfoLevel
	e.handle()
}

// Notice logs a notice log entry
func (e Entry) Notice(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = NoticeLevel
	e.handle()
}

// Noticef logs a notice log entry with formatting
func (e Entry) Noticef(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = NoticeLevel
	e.handle()
}

// Warn logs a warn log entry
func (e Entry) Warn(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = WarnLevel
	e.Handler.Handle(&e)
}

// Warnf logs a warn log entry with formatting
func (e Entry) Warnf(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = WarnLevel
	e.handle()
}

// Panic logs a panic log entry
func (e Entry) Panic(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = PanicLevel
	e.handle()
}

// Panicf logs a panic log entry with formatting
func (e Entry) Panicf(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = PanicLevel
	e.handle()
}

// Alert logs an alert log entry
func (e Entry) Alert(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = AlertLevel
	e.handle()
}

// Alertf logs an alert log entry with formatting
func (e Entry) Alertf(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = AlertLevel
	e.handle()
}

// Fatal logs a fatal log entry
func (e Entry) Fatal(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = FatalLevel
	e.handle()
}

// Fatalf logs a fatal log entry with formatting
func (e Entry) Fatalf(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = FatalLevel
	e.handle()
}

// Error logs an error log entry
func (e Entry) Error(v ...interface{}) {
	e.Message = fmt.Sprint(v...)
	e.Level = ErrorLevel
	e.handle()
}

// Errorf logs an error log entry with formatting
func (e Entry) Errorf(s string, v ...interface{}) {
	e.Message = fmt.Sprintf(s, v...)
	e.Level = ErrorLevel
	e.handle()
}

func (e Entry) handle() {
	e.Timestamp = time.Now()

	if e.Handler != nil {
		e.Handler.Handle(&e)
	}

	if e.Exit != nil {
		switch e.Level {
		case PanicLevel:
			e.Exit(1)
		case FatalLevel:
			e.Exit(1)
		}
	}
}
