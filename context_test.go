package log_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
)

var _ = Describe("Context", func() {
	It("sets the entry", func() {
		w := log.New(&log.Config{})

		ctx := log.SetContext(context.Background(), w)
		Expect(log.GetContext(ctx)).To(Equal(w))
	})

	It("gets the entry", func() {
		ctx := context.Background()
		e := log.GetContext(ctx)
		Expect(e).NotTo(BeNil())
	})

	Describe("LogEntryCtxKey", func() {
		It("formats as string", func() {
			Expect(log.LoggerCtxKey.String()).To(Equal("log: context value logger"))
		})
	})
})
