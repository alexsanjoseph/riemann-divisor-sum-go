package riemann_test

import (
	"fmt"
	"math"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	"github.com/dustin/go-humanize"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CounterExample Search", func() {

	It("should return witness value", func() {
		output := riemann.WitnessValue(10080)
		Expect(math.Abs(output-1.755814) < 1e-5).To(BeTrue())
	})

	It("should fail if no witnesses", func() {
		_, err := riemann.Search(6000, 5041)
		Expect(err).To(HaveOccurred())
	})

	It("should search successfully", func() {
		output, err := riemann.Search(10000, 5040)
		if err != nil {
			Fail("error should be nil")
		}
		Expect(output).To(Equal(5040))
	})

	It("should find best witness successfully", func() {
		count_till := 1000000
		output, witnessVal := riemann.BestWitness(count_till, 5041)
		fmt.Println("\nCurrent Best till", humanize.Comma(int64(count_till)), "is", output, "at value", witnessVal)

		Expect(output).To(Equal(10080))
	})
})
