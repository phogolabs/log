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

			Expect(fields).To(HaveKey("source"))
			Expect(fields).To(HaveKeyWithValue("error", "ginkgo: oh no!"))
			Expect(fields).To(HaveKeyWithValue("code", 1))
			Expect(fields).To(HaveKeyWithValue("types", "permanent"))
		})
	})

	Context("when the error is causer", func() {
		It("returns the fields", func() {
			chain := perrors.Wrap(err, "ginkgo")

			fields := log.FieldsOfError(chain)
			Expect(fields).To(HaveLen(2))

			Expect(fields).To(HaveKey("source"))
			Expect(fields).To(HaveKeyWithValue("error", "ginkgo: oh no!"))
		})
	})
})
