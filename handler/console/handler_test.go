package console_test

import (
	"bytes"
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
			},
			Handler: handler,
		}

		handler.Handle(entry)

		Expect(buffer.String()).To(ContainSubstring("ginkgo"))
	})
})
