package console_test

import (
	"bytes"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/handler/console"
)

var _ = Describe("Handler", func() {
	var (
		buffer  *bytes.Buffer
		handler *console.Handler
	)

	BeforeEach(func() {
		buffer = &bytes.Buffer{}
		handler = console.New(buffer)
	})

	It("writes toe the buffer", func() {
		entry := &log.Entry{
			Message:   "hello",
			Timestamp: time.Now(),
			Level:     log.InfoLevel,
			Fields: []log.Field{
				log.F("app", "ginkgo"),
				log.F("int", int(1)),
				log.F("int8", int8(1)),
				log.F("int16", int16(1)),
				log.F("int32", int32(1)),
				log.F("int64", int64(1)),
				log.F("uint", uint(1)),
				log.F("uint8", uint8(1)),
				log.F("uint16", uint16(1)),
				log.F("uint32", uint32(1)),
				log.F("uint64", uint64(1)),
				log.F("float32", float32(1)),
				log.F("float64", float64(1)),
				log.F("bool", true),
				log.F("url", &url.URL{}),
			},
			Handler: handler,
		}

		handler.Handle(entry)

		Expect(buffer.String()).To(ContainSubstring("ginkgo"))
	})
})
