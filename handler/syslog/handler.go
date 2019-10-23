package syslog

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/go-playground/ansi"
	"github.com/phogolabs/log"
)

const (
	space   = byte(' ')
	equals  = byte('=')
	newLine = byte('\n')
	base10  = 10
	pattern = "%v"
)

var _ log.Handler = &Handler{}

//go:generate counterfeiter -fake-name Syslogger -o ../../fake/syslogger.go . Logger

// Logger writer
type Logger interface {
	Alert(m string) error
	Crit(m string) error
	Debug(m string) error
	Emerg(m string) error
	Err(m string) error
	Info(m string) error
	Notice(m string) error
	Warning(m string) error
	Write(p []byte) (n int, err error)
}

// Config is the configuration
type Config struct {
	Logger          Logger
	TimestampFormat string
}

// Handler is an instance of the console logger
type Handler struct {
	logger          Logger
	timestampFormat string
}

// NewConfig creates a handler with a config
func NewConfig(config *Config) *Handler {
	return &Handler{
		timestampFormat: config.TimestampFormat,
		logger:          config.Logger,
	}
}

// New creates a new console handler
func New(logger Logger) *Handler {
	config := &Config{
		Logger:          logger,
		TimestampFormat: "2006-01-02 15:04:05.000000000Z07:00",
	}

	return NewConfig(config)
}

// Handle handles the log entry
func (c *Handler) Handle(e *log.Entry) {
	var (
		line  []byte
		color ansi.EscSeq
		level string
	)

	line = append(line, e.Timestamp.Format(c.timestampFormat)...)
	line = append(line, space)

	level = e.Level.String()

	for i := 0; i < 6-len(level); i++ {
		line = append(line, space)
	}

	line = append(line, level...)
	line = append(line, ansi.Reset...)
	line = append(line, space)
	line = append(line, e.Message...)

	for _, key := range c.keys(e) {
		value := e.Fields[key]
		line = append(line, space)

		if len(color) > 0 {
			line = append(line, color...)
		}

		line = append(line, key...)
		line = append(line, ansi.Reset...)
		line = append(line, equals)

		switch t := value.(type) {
		case string:
			line = append(line, t...)
		case int:
			line = strconv.AppendInt(line, int64(t), base10)
		case int8:
			line = strconv.AppendInt(line, int64(t), base10)
		case int16:
			line = strconv.AppendInt(line, int64(t), base10)
		case int32:
			line = strconv.AppendInt(line, int64(t), base10)
		case int64:
			line = strconv.AppendInt(line, t, base10)
		case uint:
			line = strconv.AppendUint(line, uint64(t), base10)
		case uint8:
			line = strconv.AppendUint(line, uint64(t), base10)
		case uint16:
			line = strconv.AppendUint(line, uint64(t), base10)
		case uint32:
			line = strconv.AppendUint(line, uint64(t), base10)
		case uint64:
			line = strconv.AppendUint(line, t, base10)
		case float32:
			line = strconv.AppendFloat(line, float64(t), 'f', -1, 32)
		case float64:
			line = strconv.AppendFloat(line, t, 'f', -1, 64)
		case bool:
			line = strconv.AppendBool(line, t)
		default:
			line = append(line, fmt.Sprintf(pattern, value)...)
		}
	}

	line = append(line, newLine)

	text := string(line)

	switch e.Level {
	case log.DebugLevel:
		c.logger.Debug(text)
	case log.InfoLevel:
		c.logger.Info(text)
	case log.NoticeLevel:
		c.logger.Notice(text)
	case log.WarnLevel:
		c.logger.Warning(text)
	case log.ErrorLevel:
		c.logger.Err(text)
	case log.PanicLevel, log.AlertLevel:
		c.logger.Alert(text)
	case log.FatalLevel:
		c.logger.Crit(text)
	}
}

func (c *Handler) keys(e *log.Entry) []string {
	keys := make([]string, 0, len(e.Fields))

	for key := range e.Fields {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}
