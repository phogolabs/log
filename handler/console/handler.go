package console

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
	"strconv"

	"github.com/go-playground/ansi"
	"github.com/phogolabs/log"
)

const (
	space   = byte(' ')
	equals  = byte('=')
	newLine = byte('\n')
	base10  = 10
	v       = "%v"
)

var _ log.Handler = &Handler{}

type syslog interface {
	Alert(m string) error
	Crit(m string) error
	Debug(m string) error
	Emerg(m string) error
	Err(m string) error
	Info(m string) error
	Notice(m string) error
	Warning(m string) error
}

// Handler is an instance of the console logger
type Handler struct {
	colors          [8]ansi.EscSeq
	writer          io.Writer
	timestampFormat string
	displayColor    bool
}

// Colors mapping.
var defaultColors = [...]ansi.EscSeq{
	log.DebugLevel:  ansi.Green,
	log.InfoLevel:   ansi.Blue,
	log.NoticeLevel: ansi.LightCyan,
	log.WarnLevel:   ansi.LightYellow,
	log.ErrorLevel:  ansi.LightRed,
	log.PanicLevel:  ansi.Red,
	log.AlertLevel:  ansi.Red + ansi.Underline,
	log.FatalLevel:  ansi.Red + ansi.Underline + ansi.Blink,
}

// New creates a new console handler
func New(writer io.Writer) *Handler {
	return &Handler{
		colors:          defaultColors,
		writer:          writer,
		timestampFormat: "2006-01-02 15:04:05.000000000Z07:00",
		displayColor:    true,
	}
}

// SetDisplayColor tells Handler to output in color or not
// Default is : true
func (c *Handler) SetDisplayColor(value bool) {
	c.displayColor = value
}

// SetTimestampFormat sets Handler's timestamp output format
// Default is : "2006-01-02T15:04:05.000000000Z07:00"
func (c *Handler) SetTimestampFormat(format string) {
	c.timestampFormat = format
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

	if c.displayColor {
		color = c.colors[e.Level]
		line = append(line, color...)
	}

	level = e.Level.String()

	for i := 0; i < 6-len(level); i++ {
		line = append(line, space)
	}

	line = append(line, level...)
	line = append(line, ansi.Reset...)
	line = append(line, space)
	line = append(line, e.Message...)

	for _, f := range e.Fields {
		line = append(line, space)

		if len(color) > 0 {
			line = append(line, color...)
		}

		line = append(line, f.Key...)
		line = append(line, ansi.Reset...)
		line = append(line, equals)

		switch t := f.Value.(type) {
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
			line = append(line, fmt.Sprintf(v, f.Value)...)
		}
	}

	line = append(line, newLine)

	if logger, ok := c.writer.(syslog); ok {
		text := string(line)

		switch e.Level {
		case log.DebugLevel:
			logger.Debug(text)
		case log.InfoLevel:
			logger.Info(text)
		case log.NoticeLevel:
			logger.Notice(text)
		case log.WarnLevel:
			logger.Warning(text)
		case log.ErrorLevel:
			logger.Err(text)
		case log.PanicLevel, log.AlertLevel:
			logger.Alert(text)
		case log.FatalLevel:
			logger.Crit(text)
		}

		return
	}

	c.writer.Write(line)
}
