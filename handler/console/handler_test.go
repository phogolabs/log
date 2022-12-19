package console_test

import (
	"bytes"
	"net/url"
	"time"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/handler/console"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	var (
		buffer  *bytes.Buffer
		handler *console.Handler
		entry   *log.Entry
	)

	BeforeEach(func() {
		buffer = &bytes.Buffer{}

		handler = console.NewConfig(&console.Config{
			Writer:          buffer,
			DisplayColor:    true,
			TimestampFormat: time.UnixDate,
		})

		entry = &log.Entry{
			Message:   "hello",
			Timestamp: time.Now(),
			Level:     log.InfoLevel,
			Fields: log.Map{
				"app":     "ginkgo",
				"int":     int(1),
				"int8":    int8(1),
				"int16":   int16(1),
				"int32":   int32(1),
				"int64":   int64(1),
				"uint":    uint(1),
				"uint8":   uint8(1),
				"uint16":  uint16(1),
				"uint32":  uint32(1),
				"uint64":  uint64(1),
				"float32": float32(1),
				"float64": float64(1),
				"bool":    true,
				"url":     &url.URL{},
			},
		}
	})

	It("writes to the buffer", func() {
		handler.Handle(entry)

		Expect(buffer.String()).To(ContainSubstring("ginkgo"))
	})
})
