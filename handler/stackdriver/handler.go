package stackdriver

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/phogolabs/log"
	logging "google.golang.org/api/logging/v2"
)

// Config is the configuration of the handler
type Config struct {
	Writer io.Writer
}

var _ log.Handler = &Handler{}

// Handler implementation.
type Handler struct {
	writer io.Writer
}

// New returns the default implementation of a Client.
func New(config *Config) *Handler {
	return &Handler{
		writer: config.Writer,
	}
}

// Handle handles the log entry
func (h *Handler) Handle(e *log.Entry) {
	entry := logging.LogEntry{
		Timestamp:   e.Timestamp.Format(time.RFC3339),
		TextPayload: e.Message,
		Severity:    h.severity(e.Level),
	}

	if trace, ok := e.Fields["trace"]; ok {
		entry.Trace = trace
	}

	if data, err := json.Marshal(&e.Fields); err == nil {
		entry.JsonPayload = data
	}

	if data, err := json.Marshal(&entry); err == nil {
		fmt.Fprintln(h.writer, string(data))
	}
}

func (h *Handler) severity(level log.Level) string {
	switch level {
	case log.InfoLevel:
		return "INFO"
	case log.DebugLevel:
		return "DEBUG"
	case log.NoticeLevel:
		return "NOTICE"
	case log.WarnLevel:
		return "WARNING"
	case log.ErrorLevel:
		return "ERROR"
	case log.PanicLevel:
		return "CRITICAL"
	case log.AlertLevel:
		return "ALERT"
	case log.FatalLevel:
		return "FATAL"
	default:
		return "DEFAULT"
	}
}
