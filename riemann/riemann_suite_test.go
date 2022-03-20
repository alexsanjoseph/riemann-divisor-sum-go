package riemann_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDivisor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Divisor")
}
