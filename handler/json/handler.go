package json

// The MIT License (MIT)

// Copyright (c) 2018 Go Playground

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/phogolabs/log"
)

var _ log.Handler = &Handler{}

// Handler implementation.
type Handler struct {
	formatter *Formatter
	writer    io.Writer
	m         sync.Mutex
}

// New handler.
func New(writer io.Writer) *Handler {
	formatter := NewFormatter()
	formatter.ColorFn = colorify

	return &Handler{
		writer:    writer,
		formatter: formatter,
	}
}

// Handle handles the log entry
func (h *Handler) Handle(e *log.Entry) {
	h.m.Lock()
	defer h.m.Unlock()

	if data, err := h.formatter.Marshal(e); err == nil {
		fmt.Fprintln(h.writer, string(data))
	}
}

// SetPretty enables pretty output
func (h *Handler) SetPretty(value bool) {
	h.formatter.DisablePretty = !value
}

// SetColor enables color formatting
func (h *Handler) SetColor(value bool) {
	h.formatter.DisabledColor = !value
}

func colorify(k, v interface{}) ColorFormatter {
	toString := func(n interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%v", n))
	}

	if name := toString(k); name != "LEVEL" {
		return nil
	}

	if level, err := log.ParseLevel(toString(v)); err == nil {
		switch level {
		case log.DebugLevel:
			return color.New(color.Reset)
		case log.InfoLevel:
			return color.New(color.FgWhite, color.Bold)
		case log.NoticeLevel:
			return color.New(color.FgHiWhite, color.Bold)
		case log.WarnLevel:
			return color.New(color.FgHiYellow, color.Bold)
		case log.ErrorLevel:
			return color.New(color.FgHiRed, color.Bold)
		case log.PanicLevel:
			return color.New(color.FgHiRed, color.Bold)
		case log.AlertLevel:
			return color.New(color.FgHiRed, color.Bold)
		case log.FatalLevel:
			return color.New(color.FgHiRed, color.Bold)
		}

	}

	return nil
}
