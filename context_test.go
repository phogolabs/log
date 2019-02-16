package log_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
)

var _ = Describe("Context", func() {
	It("sets the entry", func() {
		entry := log.Entry{}

		ctx := log.SetContext(context.Background(), entry)
		Expect(log.GetContext(ctx)).To(Equal(entry))
	})

	It("gets the entry", func() {
		ctx := context.Background()
		e := log.GetContext(ctx)
		Expect(e.Exit).NotTo(BeNil())
	})

	Describe("LogEntryCtxKey", func() {
		It("formats as string", func() {
			Expect(log.EntryCtxKey.String()).To(Equal("log: context value entry"))
		})
	})
})
