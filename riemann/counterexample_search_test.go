package riemann_test

import (
	"fmt"
	"math"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"

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
		output, witnessVal := riemann.BestWitness(100000, 5041)
		fmt.Println("Current Best", output, "at value", witnessVal)

		Expect(output).To(Equal(10080))
	})
})
