package console_test

import (
	"bytes"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/fake"
	"github.com/phogolabs/log/handler/console"
)

var _ = Describe("Handler", func() {
	var (
		buffer  *bytes.Buffer
		handler *console.Handler
		entry   *log.Entry
	)

	BeforeEach(func() {
		buffer = &bytes.Buffer{}
		handler = console.New(buffer)

		handler.SetDisplayColor(true)
		handler.SetTimestampFormat(time.UnixDate)

		entry = &log.Entry{
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
	})

	It("writes to the buffer", func() {
		handler.Handle(entry)

		Expect(buffer.String()).To(ContainSubstring("ginkgo"))
	})

	Context("when the writter is syslog", func() {
		var syslogger *fake.Syslogger

		BeforeEach(func() {
			syslogger = &fake.Syslogger{}
			handler = console.New(syslogger)
		})

		It("writes to the debug buffer", func() {
			entry.Level = log.DebugLevel
			handler.Handle(entry)
			Expect(syslogger.DebugCallCount()).To(Equal(1))
		})

		It("writes to the info buffer", func() {
			entry.Level = log.InfoLevel
			handler.Handle(entry)
			Expect(syslogger.InfoCallCount()).To(Equal(1))
		})

		It("writes to the notice buffer", func() {
			entry.Level = log.NoticeLevel
			handler.Handle(entry)
			Expect(syslogger.NoticeCallCount()).To(Equal(1))
		})

		It("writes to the warn buffer", func() {
			entry.Level = log.WarnLevel
			handler.Handle(entry)
			Expect(syslogger.WarningCallCount()).To(Equal(1))
		})

		It("writes to the error buffer", func() {
			entry.Level = log.ErrorLevel
			handler.Handle(entry)
			Expect(syslogger.ErrCallCount()).To(Equal(1))
		})

		It("writes to the alert buffer", func() {
			entry.Level = log.AlertLevel
			handler.Handle(entry)
			Expect(syslogger.AlertCallCount()).To(Equal(1))
		})

		It("writes to the critical buffer", func() {
			entry.Level = log.FatalLevel
			handler.Handle(entry)
			Expect(syslogger.CritCallCount()).To(Equal(1))
		})
	})
})
