package syslog_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSyslog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Syslog Suite")
}
