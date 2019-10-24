package json_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/phogolabs/log"
	"github.com/phogolabs/log/handler/json"
)

var _ = Describe("Handler", func() {

	var (
		buffer  *gbytes.Buffer
		handler *json.Handler
	)

	BeforeEach(func() {
		buffer = gbytes.NewBuffer()
		handler = json.New(buffer)
	})

	ItEncodesTheEntry := func(level log.Level) {
		It("encodes the entry as json", func() {
			entry := &log.Entry{
				Message:   "hello",
				Timestamp: time.Now(),
				Level:     level,
				Fields: log.Map{
					"app": "ginkgo",
				},
			}

			handler.Handle(entry)

			Expect(buffer).To(gbytes.Say(level.String()))
		})
	}

	ItHandlesTheLevels := func() {
		Describe("Handle", func() {
			Context("when the level is INFO", func() {
				ItEncodesTheEntry(log.InfoLevel)
			})

			Context("when the level is WARN", func() {
				ItEncodesTheEntry(log.WarnLevel)
			})

			Context("when the level is NOTICE", func() {
				ItEncodesTheEntry(log.NoticeLevel)
			})

			Context("when the level is ERROR", func() {
				ItEncodesTheEntry(log.ErrorLevel)
			})

			Context("when the level is PANIC", func() {
				ItEncodesTheEntry(log.PanicLevel)
			})

			Context("when the level is ALERT", func() {
				ItEncodesTheEntry(log.AlertLevel)
			})

			Context("when the level is FATAL", func() {
				ItEncodesTheEntry(log.FatalLevel)
			})
		})
	}

	ItHandlesTheLevels()

	Context("when the pretty and color are on", func() {
		BeforeEach(func() {
			handler = json.NewConfig(&json.Config{
				Writer:           buffer,
				DisplayColor:     true,
				PrettyFormatting: true,
			})
		})

		ItHandlesTheLevels()
	})
})
