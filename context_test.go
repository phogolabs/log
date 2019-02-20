package log_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
)

var _ = Describe("Context", func() {
	It("sets the entry", func() {
		w := log.NewWriter(&log.WriterConfig{})

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
			Expect(log.WriterCtxKey.String()).To(Equal("log: context value writer"))
		})
	})
})
