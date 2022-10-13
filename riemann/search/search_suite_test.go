package search_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/types"
	. "github.com/onsi/gomega"
)

func TestDivisor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Search Tests", types.ReporterConfig{
		SlowSpecThreshold: 100 * time.Millisecond,
	})
}
