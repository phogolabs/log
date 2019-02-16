package log_test

import (
	"encoding/json"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
)

var _ = Describe("Level", func() {
	DescribeTable("ParseLevel", func(text string) {
		level, err := log.ParseLevel(text)
		Expect(err).NotTo(HaveOccurred())
		Expect(level.String()).To(Equal(strings.ToUpper(text)))
	},
		Entry("when the level is debug", "debug"),
		Entry("when the level is info", "info"),
		Entry("when the level is notice", "notice"),
		Entry("when the level is warn", "warn"),
		Entry("when the level is error", "error"),
		Entry("when the level is panic", "panic"),
		Entry("when the level is fatal", "fatal"),
		Entry("when the level is alert", "alert"),
	)

	Context("when the level is wrong", func() {
		It("returns an error", func() {
			level, err := log.ParseLevel("verbose")
			Expect(err).To(MatchError("log: unknown level \"verbose\""))
			Expect(level).To(BeNumerically("==", 255))
		})

		It("formats the level", func() {
			level := log.Level(255)
			Expect(level.String()).To(Equal("UNKNOWN"))
		})
	})

	Describe("MarshalJSON", func() {
		It("marshals a level successfully", func() {
			data, err := json.Marshal(log.InfoLevel)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(Equal([]byte("\"INFO\"")))
		})

	})

	Describe("UnmarshalJSON", func() {
		It("unmarshals level successfully", func() {
			var (
				level log.Level
				data  = []byte("\"INFO\"")
			)

			Expect(json.Unmarshal(data, &level)).To(Succeed())
			Expect(level).To(Equal(log.InfoLevel))
		})

		Context("when the level is wrong", func() {
			It("returns an error", func() {
				var (
					level log.Level
					data  = []byte("\"verbose\"")
				)

				Expect(json.Unmarshal(data, &level)).To(MatchError("log: unknown level \"verbose\""))
			})
		})
	})
})
