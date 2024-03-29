package log_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/fake"
)

var _ = Describe("Log", func() {
	var handler *fake.Handler

	BeforeEach(func() {
		handler = &fake.Handler{}

		log.SetLevel(log.DebugLevel)
		log.SetHandler(handler)
		log.SetExitFunc(func(code int) {
			Expect(code).To(Equal(1))
		})

		log.SetDefaultFields(log.Map{})
	})

	Describe("SetDefaultFields", func() {
		It("sets the default fields", func() {
			log.SetDefaultFields(log.Map{"app": "ginkgo"})

			e := log.WithField("version", "beta").(log.Logger)

			Expect(e.Fields()).To(HaveLen(2))
			Expect(e.Fields()).To(HaveKeyWithValue("app", "ginkgo"))
			Expect(e.Fields()).To(HaveKeyWithValue("version", "beta"))
		})
	})

	Describe("SetDefaultFieldsWithMap", func() {
		It("sets the default fields", func() {
			log.SetDefaultFields(log.Map{"app": "ginkgo"})

			e := log.WithField("version", "beta").(log.Logger)

			Expect(e.Fields()).To(HaveLen(2))
			Expect(e.Fields()).To(HaveKeyWithValue("app", "ginkgo"))
			Expect(e.Fields()).To(HaveKeyWithValue("version", "beta"))
		})
	})

	Describe("SetLevel", func() {
		It("sets the level", func() {
			log.SetLevel(log.ErrorLevel)
			log.Info("oh no!")

			Expect(handler.HandleCallCount()).To(Equal(0))
		})
	})

	Describe("WithField", func() {
		It("returns a new entry", func() {
			e := log.WithField("address", "service-api").(log.Logger)
			Expect(e.Fields()).To(HaveLen(1))
			Expect(e.Fields()).To(HaveKeyWithValue("address", "service-api"))
		})
	})

	Describe("WithFields", func() {
		It("returns a new entry", func() {
			e := log.WithField("address", "service-api").(log.Logger)
			Expect(e.Fields()).To(HaveLen(1))
			Expect(e.Fields()).To(HaveKeyWithValue("address", "service-api"))
		})

		Context("when is a map", func() {
			It("returns a new entry", func() {
				fields := log.Map{
					"address": "service-api",
				}

				e := log.WithFields(fields).(log.Logger)
				Expect(e.Fields()).To(HaveLen(1))
				Expect(e.Fields()).To(HaveKeyWithValue("address", "service-api"))
			})
		})
	})

	Describe("WithError", func() {
		It("returns a new entry", func() {
			err := fmt.Errorf("oh no")

			e := log.WithError(err).(log.Logger)
			Expect(e.Fields()).To(HaveLen(2))

			fields := e.Fields()
			Expect(fields).To(HaveKey("error_stack"))
			Expect(fields).To(HaveKeyWithValue("error_cause", "oh no"))
		})
	})

	DescribeOperation := func(level log.Level) {
		op := func(v ...interface{}) {
			switch level {
			case log.DebugLevel:
				log.Debug(v...)
			case log.InfoLevel:
				log.Info(v...)
			case log.NoticeLevel:
				log.Notice(v...)
			case log.WarnLevel:
				log.Warn(v...)
			case log.PanicLevel:
				log.Panic(v...)
			case log.AlertLevel:
				log.Alert(v...)
			case log.FatalLevel:
				log.Fatal(v...)
			case log.ErrorLevel:
				log.Error(v...)
			}
		}

		opf := func(msg string, v ...interface{}) {
			switch level {
			case log.DebugLevel:
				log.Debugf(msg, v...)
			case log.InfoLevel:
				log.Infof(msg, v...)
			case log.NoticeLevel:
				log.Noticef(msg, v...)
			case log.WarnLevel:
				log.Warnf(msg, v...)
			case log.PanicLevel:
				log.Panicf(msg, v...)
			case log.AlertLevel:
				log.Alertf(msg, v...)
			case log.FatalLevel:
				log.Fatalf(msg, v...)
			case log.ErrorLevel:
				log.Errorf(msg, v...)
			}
		}

		Describe(level.String(), func() {
			It("writes the message", func() {
				op("hello")

				Expect(handler.HandleCallCount()).To(Equal(1))

				e := handler.HandleArgsForCall(0)
				Expect(e.Level).To(Equal(level))
				Expect(e.Message).To(Equal("hello"))
			})

			Context("when the message is formatted", func() {
				It("returns a new entry", func() {
					opf("hello, %v", "jack")

					Expect(handler.HandleCallCount()).To(Equal(1))

					e := handler.HandleArgsForCall(0)
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

var _ = Describe("DefaultHandler", func() {
	var handler *log.DefaultHandler

	BeforeEach(func() {
		handler = &log.DefaultHandler{}
	})

	It("logs entry", func() {
		entry := &log.Entry{}
		handler.Handle(entry)
	})
})
