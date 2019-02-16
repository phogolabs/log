package json_test

import (
	"encoding/json"
	"io"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Json Suite")
}

func unmarshal(reader io.Reader) map[string]interface{} {
	m := make(map[string]interface{})
	Expect(json.NewDecoder(reader).Decode(&m)).To(Succeed())
	return m
}
