package protocol_integration_tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProtocolIntegrationTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProtocolIntegrationTests Suite")
}
