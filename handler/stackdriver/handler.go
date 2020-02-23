package stackdriver

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/phogolabs/log"
)

// LogEntry represents the Google Cloud Log Entry
type LogEntry struct {
	// TextPayload: The log entry payload, represented as a Unicode string
	// (UTF-8).
	TextPayload string `json:"textPayload,omitempty"`

	// Timestamp: Optional. The time the event described by the log entry
	// occurred. This time is used to compute the log entry's age and to
	// enforce the logs retention period. If this field is omitted in a new
	// log entry, then Logging assigns it the current time. Timestamps have
	// nanosecond accuracy, but trailing zeros in the fractional seconds
	// might be omitted when the timestamp is displayed.Incoming log entries
	// should have timestamps that are no more than the logs retention
	// period in the past, and no more than 24 hours in the future. Log
	// entries outside those time boundaries will not be available when
	// calling entries.list, but those log entries can still be exported
	// with LogSinks.
	Timestamp string `json:"timestamp,omitempty"`

	// Trace: Optional. Resource name of the trace associated with the log
	// entry, if any. If it contains a relative resource name, the name is
	// assumed to be relative to //tracing.googleapis.com. Example:
	// projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824
	Trace string `json:"trace,omitempty"`

	// Severity: Optional. The severity of the log entry. The default value
	// is LogSeverity.DEFAULT.
	//
	// Possible values:
	//   "DEFAULT" - (0) The log entry has no assigned severity level.
	//   "DEBUG" - (100) Debug or trace information.
	//   "INFO" - (200) Routine information, such as ongoing status or
	// performance.
	//   "NOTICE" - (300) Normal but significant events, such as start up,
	// shut down, or a configuration change.
	//   "WARNING" - (400) Warning events might cause problems.
	//   "ERROR" - (500) Error events are likely to cause problems.
	//   "CRITICAL" - (600) Critical events cause more severe problems or
	// outages.
	//   "ALERT" - (700) A person must take an action immediately.
	//   "EMERGENCY" - (800) One or more systems are unusable.
	Severity string `json:"severity,omitempty"`

	// Fields represents the associated fields
	Fields log.Map `json:"fields,omitempty"`
}

// Config is the configuration of the handler
type Config struct {
	ProjectID string
	Writer    io.Writer
}

var _ log.Handler = &Handler{}

// Handler implementation.
type Handler struct {
	projectID string
	writer    io.Writer
}

// New returns the default implementation of a Client.
func New(config *Config) *Handler {
	return &Handler{
		projectID: config.ProjectID,
		writer:    config.Writer,
	}
}

// Handle handles the log entry
func (h *Handler) Handle(e *log.Entry) {
	entry := LogEntry{
		Timestamp:   e.Timestamp.Format(time.RFC3339),
		TextPayload: e.Message,
		Fields:      e.Fields,
		Severity:    h.severity(e.Level),
	}

	if h.projectID != "" {
		if trace, ok := e.Fields["trace_context"].(string); ok {
			if parts := strings.Split(trace, "/"); len(parts) > 0 {
				if head := parts[0]; len(head) > 0 {
					entry.Trace = fmt.Sprintf("projects/%s/traces/%s", h.projectID, head)
				}
			}
		}
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
