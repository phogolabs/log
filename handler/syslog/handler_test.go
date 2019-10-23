package syslog_test

import (
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/fake"
	"github.com/phogolabs/log/handler/syslog"
)

var _ = Describe("Handler", func() {
	var (
		handler   *syslog.Handler
		entry     *log.Entry
		syslogger *fake.Syslogger
	)

	BeforeEach(func() {
		syslogger = &fake.Syslogger{}

		handler = syslog.New(syslogger)

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
