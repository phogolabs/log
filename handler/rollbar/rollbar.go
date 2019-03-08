// Package rollbar implements a JSON handler.
package rollbar

import (
	log "github.com/phogolabs/log"
	rollbar "github.com/rollbar/rollbar-go"
)

// Client that connects to rollbar
type Client interface {
	MessageWithExtras(level string, msg string, extras map[string]interface{})
}

// Config is the configuration of the handler
type Config struct {
	Token       string
	Environment string
	CodeVersion string
	ServerHost  string
	ServerRoot  string
}

var _ log.Handler = &Handler{}

// Handler implementation.
type Handler struct {
	Client Client
}

// New returns the default implementation of a Client.
func New(config *Config) *Handler {
	return &Handler{
		Client: rollbar.NewAsync(
			config.Token,
			config.Environment,
			config.CodeVersion,
			config.ServerHost,
			config.ServerRoot,
		),
	}
}

// Handle handles the log entry
func (h *Handler) Handle(e *log.Entry) {
	var (
		level  string
		extras = make(map[string]interface{}, len(e.Fields))
	)

	switch e.Level {
	case log.DebugLevel:
		level = rollbar.DEBUG
	case log.InfoLevel:
		level = rollbar.INFO
	case log.NoticeLevel, log.WarnLevel:
		level = rollbar.WARN
	case log.ErrorLevel:
		level = rollbar.ERR
	case log.PanicLevel, log.AlertLevel, log.FatalLevel:
		level = rollbar.CRIT
	}

	for _, fielder := range e.Fields {
		for _, field := range fielder.Fields() {
			extras[field.Key] = field.Value
		}
	}

	h.Client.MessageWithExtras(level, e.Message, extras)
}
