package json_test

import (
	"bytes"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/log"
	"github.com/phogolabs/log/handler/json"
)

var _ = Describe("Handler", func() {

	var (
		buffer  *bytes.Buffer
		handler *json.Handler
	)

	BeforeEach(func() {
		buffer = &bytes.Buffer{}
		handler = json.New(buffer)
	})

	It("encodes the entry as json", func() {
		entry := &log.Entry{
			Message:   "hello",
			Timestamp: time.Now(),
			Level:     log.InfoLevel,
			Fields: log.FieldMap{
				"app": "ginkgo",
			},
		}

		handler.Handle(entry)

		m := unmarshal(buffer)
		Expect(m).To(HaveKeyWithValue("message", "hello"))
		Expect(m).To(HaveKeyWithValue("level", "INFO"))
	})
})
