package log_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/phogolabs/log"

	gerrors "github.com/go-playground/errors"
	perrors "github.com/pkg/errors"
)

var _ = Describe("Error", func() {
	err := fmt.Errorf("oh no!")

	Context("when the error is chained", func() {
		It("returns the fields", func() {
			chain := gerrors.Wrap(err, "ginkgo")
			chain.AddTag("code", 1)
			chain.AddTypes("permanent")

			fields := log.FieldsOfError(chain)
			Expect(fields).To(HaveLen(4))

			Expect(fields[0].Key).To(Equal("source"))

			Expect(fields[1].Key).To(Equal("error"))
			Expect(fields[1].Value).To(Equal("ginkgo: oh no!"))

			Expect(fields[2].Key).To(Equal("code"))
			Expect(fields[2].Value).To(Equal(1))

			Expect(fields[3].Key).To(Equal("types"))
			Expect(fields[3].Value).To(Equal("permanent"))
		})
	})

	Context("when the error is causer", func() {
		It("returns the fields", func() {
			chain := perrors.Wrap(err, "ginkgo")

			fields := log.FieldsOfError(chain)
			Expect(fields).To(HaveLen(2))

			Expect(fields[0].Key).To(Equal("error"))
			Expect(fields[0].Value).To(Equal("ginkgo: oh no!"))

			Expect(fields[1].Key).To(Equal("source"))
		})
	})
})
