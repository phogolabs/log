package log

import (
	"encoding/json"
	"log"
)

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
