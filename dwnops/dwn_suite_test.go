package dwnops_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDwn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dwn Suite")
}
