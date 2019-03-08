package log_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/fake"
)

var _ = Describe("CompositeHandler", func() {
	var (
		composite log.CompositeHandler
		handler   *fake.Handler
	)

	BeforeEach(func() {
		handler = &fake.Handler{}
		composite = log.CompositeHandler{handler}
	})

	It("delegates to underlying handlers successfully", func() {
		entry := log.Entry{}
		composite.Handle(&entry)

		Expect(handler.HandleCallCount()).To(Equal(1))
		Expect(handler.HandleArgsForCall(0)).To(Equal(&entry))
	})
})

var _ = Describe("LevelHandler", func() {
	var (
		composite *log.LevelHandler
		handler   *fake.Handler
	)

	BeforeEach(func() {
		handler = &fake.Handler{}
		composite = &log.LevelHandler{
			Level:   log.InfoLevel,
			Handler: handler,
		}
	})

	Context("when the entry's level < Info", func() {
		It("does not delegate to the underlying handler", func() {
			entry := log.Entry{Level: log.DebugLevel}

			composite.Handle(&entry)

			Expect(handler.HandleCallCount()).To(Equal(0))
		})
	})

	Context("when the entry's level >= Info", func() {
		It("does not delegate to the underlying handler", func() {
			entry := log.Entry{Level: log.InfoLevel}

			composite.Handle(&entry)

			Expect(handler.HandleCallCount()).To(Equal(1))
			Expect(handler.HandleArgsForCall(0)).To(Equal(&entry))
		})
	})
})

var _ = Describe("Writer", func() {
	var (
		entry   log.Entry
		writer  log.Writer
		handler *fake.Handler
	)

	BeforeEach(func() {
		handler = &fake.Handler{}

		config := log.WriterConfig{
			Handler: handler,
			Exit: func(code int) {
				Expect(code).To(Equal(1))
			},
		}

		writer = log.NewWriter(&config)
	})

	Describe("WithField", func() {
		It("returns a new entry", func() {
			e := writer.WithField("app", "service-api").(log.Writer)
			Expect(e).NotTo(Equal(entry))
			Expect(e.Entry().Fields).To(HaveLen(1))
			Expect(e.Entry().Fields).To(ContainElement(log.F("app", "service-api")))
		})
	})

	Describe("WithFields", func() {
		It("returns a new entry", func() {
			e := writer.WithFields(log.F("app", "service-api")).(log.Writer)
			Expect(e).NotTo(Equal(entry))
			Expect(e.Entry().Fields).To(HaveLen(1))
			Expect(e.Entry().Fields).To(ContainElement(log.F("app", "service-api")))
		})

		Context("when is a map", func() {
			It("returns a new entry", func() {
				fields := log.M{
					"app": "service-api",
				}

				e := writer.WithFields(fields).(log.Writer)
				Expect(e).NotTo(Equal(entry))
				Expect(e.Entry().Fields).To(HaveLen(1))
				Expect(e.Entry().Fields[0]).To(HaveKeyWithValue("app", "service-api"))
			})
		})
	})

	Describe("WithError", func() {
		It("returns a new entry", func() {
			err := fmt.Errorf("oh no!")
			e := writer.WithError(err).(log.Writer)
			Expect(e).NotTo(Equal(entry))
			Expect(e.Entry().Fields[0]).To(HaveKeyWithValue("error", "oh no!"))
		})
	})

	DescribeOperation := func(level log.Level) {
		op := func(e log.Writer, v ...interface{}) {
			switch level {
			case log.DebugLevel:
				e.Debug(v...)
			case log.InfoLevel:
				e.Info(v...)
			case log.NoticeLevel:
				e.Notice(v...)
			case log.WarnLevel:
				e.Warn(v...)
			case log.PanicLevel:
				e.Panic(v...)
			case log.AlertLevel:
				e.Alert(v...)
			case log.FatalLevel:
				e.Fatal(v...)
			case log.ErrorLevel:
				e.Error(v...)
			}
		}

		opf := func(e log.Writer, msg string, v ...interface{}) {
			switch level {
			case log.DebugLevel:
				e.Debugf(msg, v...)
			case log.InfoLevel:
				e.Infof(msg, v...)
			case log.NoticeLevel:
				e.Noticef(msg, v...)
			case log.WarnLevel:
				e.Warnf(msg, v...)
			case log.PanicLevel:
				e.Panicf(msg, v...)
			case log.AlertLevel:
				e.Alertf(msg, v...)
			case log.FatalLevel:
				e.Fatalf(msg, v...)
			case log.ErrorLevel:
				e.Errorf(msg, v...)
			}
		}

		Describe(level.String(), func() {
			It("writes the message", func() {
				op(writer, "hello")

				Expect(handler.HandleCallCount()).To(Equal(1))

				e := handler.HandleArgsForCall(0)
				Expect(e).NotTo(Equal(&entry))
				Expect(e.Level).To(Equal(level))
				Expect(e.Message).To(Equal("hello"))
			})

			Context("when the message is formatted", func() {
				It("returns a new entry", func() {
					opf(writer, "hello, %v", "jack")

					Expect(handler.HandleCallCount()).To(Equal(1))

					e := handler.HandleArgsForCall(0)
					Expect(e).NotTo(Equal(&entry))
					Expect(e.Level).To(Equal(level))
					Expect(e.Message).To(Equal("hello, jack"))
				})
			})
		})
	}

	DescribeOperation(log.DebugLevel)
	DescribeOperation(log.InfoLevel)
	DescribeOperation(log.NoticeLevel)
	DescribeOperation(log.WarnLevel)
	DescribeOperation(log.PanicLevel)
	DescribeOperation(log.AlertLevel)
	DescribeOperation(log.FatalLevel)
	DescribeOperation(log.ErrorLevel)
})
