package rollbar_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
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

	Describe("New", func() {
		It("creates the logger successfully", func() {
			config := &rollbar.Config{}
			Expect(rollbar.New(config)).NotTo(BeNil())
		})
	})

	DescribeTable("sends the log", func(llevel log.Level, rlevel string) {
		entry := &log.Entry{
			Message:   "hello",
			Timestamp: time.Now(),
			Level:     llevel,
			Fields: log.FieldMap{
				"app": "ginkgo",
			},
		}

		handler.Handle(entry)

		Expect(client.MessageWithExtrasCallCount()).To(Equal(1))

		level, msg, extras := client.MessageWithExtrasArgsForCall(0)
		Expect(level).To(Equal(rlevel))
		Expect(msg).To(Equal("hello"))
		Expect(extras).To(HaveKeyWithValue("app", "ginkgo"))
	},
		Entry("debug", log.DebugLevel, "debug"),
		Entry("info", log.InfoLevel, "info"),
		Entry("warn", log.WarnLevel, "warning"),
		Entry("panic", log.PanicLevel, "critical"),
		Entry("error", log.ErrorLevel, "error"),
	)
})
