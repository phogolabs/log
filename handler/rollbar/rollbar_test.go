package rollbar_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	log "github.com/phogolabs/log"
	"github.com/phogolabs/log/fake"
	"github.com/phogolabs/log/handler/rollbar"
)

var _ = Describe("Rollbar", func() {
	var (
		client  *fake.Client
		handler *rollbar.Handler
	)

	BeforeEach(func() {
		client = &fake.Client{}

		handler = &rollbar.Handler{
			Client: client,
		}
	})

	It("sends the log to rollbar", func() {
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

		Expect(client.MessageWithExtrasCallCount()).To(Equal(1))

		level, msg, extras := client.MessageWithExtrasArgsForCall(0)
		Expect(level).To(Equal("info"))
		Expect(msg).To(Equal("hello"))
		Expect(extras).To(HaveKeyWithValue("app", "ginkgo"))
	})
})
