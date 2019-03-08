package log

import (
	"encoding/json"
	"os"
	"sort"
	"time"
)

// Entry defines a single log entry
type Entry struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Fields    []Fielder `json:"fields,omitempty"`
	Level     Level     `json:"level"`
}

// Writer of the log
type Writer interface {
	// WithField returns a new log entry with the supplied field.
	WithField(key string, value interface{}) Writer

	// WithFields returns a new log entry with the supplied fields appended
	WithFields(fields ...Fielder) Writer

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

	// Entry returns the entry
	Entry() Entry
}

// standard logger
var std = writer{
	exit: os.Exit,
	handler: &LevelHandler{
		Level:   DebugLevel,
		Handler: &DefaultHandler{},
	},
}

// SetExitFn sets the exit function. default: os.Exit
func SetExitFn(fn ExitFunc) {
	std.exit = fn
}

// SetLevel sets the level handler
func SetLevel(level Level) {
	leveled := std.handler.(*LevelHandler)
	leveled.Level = level
}

// SetHandler sets the handler
func SetHandler(handler Handler) {
	leveled := std.handler.(*LevelHandler)
	leveled.Handler = handler
}

// SetDefaultFields sets the default fields
func SetDefaultFields(fields ...Fielder) {
	std.entry.Fields = fields
}

// WithField returns a new log entry with the supplied field.
func WithField(key string, value interface{}) Writer {
	return std.WithField(key, value)
}

// WithFields returns a new log entry with the supplied fields appended
func WithFields(fields ...Fielder) Writer {
	return std.WithFields(fields...)
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

// ExitFunc is a function called on Panic or Fatal level
type ExitFunc func(code int)

// Fielder returns the fields
type Fielder interface {
	// Fields returns the fields
	Fields() []Field
}

// Fields map
type Fields = M

var _ Fielder = Field{}

// Field is a single Field key and value
type Field struct {
	Key   string
	Value interface{}
}

// Fields returns the fields
func (f Field) Fields() []Field {
	return []Field{f}
}

// MarshalJSON marshals the field
func (f Field) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		f.Key: f.Value,
	}

	return json.Marshal(&m)
}

// F creates a new Field using the supplied key + value.
// it is shorthand for defining field manually
func F(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

var _ Fielder = M{}

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

// FieldSorter sorts a slice of Fields to be sorted.
type FieldSorter struct {
	Fields []Field
}

// Len is part of sort.Interface.
func (s *FieldSorter) Len() int {
	return len(s.Fields)
}

// Swap is part of sort.Interface.
func (s *FieldSorter) Swap(i, j int) {
	s.Fields[i], s.Fields[j] = s.Fields[j], s.Fields[i]
}

// Less is part of sort.Interface.
func (s *FieldSorter) Less(i, j int) bool {
	return s.Fields[i].Key < s.Fields[j].Key
}

// SortFields sorts the fields
func SortFields(fields []Field) {
	sorter := &FieldSorter{Fields: fields}
	sort.Sort(sorter)
}
