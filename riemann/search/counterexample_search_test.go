package search_test

import (
	"fmt"
	"math"
	"math/big"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/search"
	"github.com/dustin/go-humanize"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CounterExample Search", func() {

	It("should return witness value", func() {
		output := search.WitnessValue(10080, -1)
		Expect(math.Abs(output-1.755814) < 1e-5).To(BeTrue())
	})

	It("should return witness value if precomputed sum is provided", func() {
		output := search.WitnessValue(10080, 1)
		Expect(math.Abs(output-(1/22389.61097)) < 1e-5).To(BeTrue())
	})

	It("should fail if no witnesses", func() {
		_, err := search.Search(6000, 5041)
		Expect(err).To(HaveOccurred())
	})

	It("should panic if asked to find witnesses for cases where DivisorSum cannot be found", func() {
		Expect(func() { search.WitnessValue(-1, -1) }).To(PanicWith("Error calculating DivisorSum for -1"))
	})

	It("should search successfully", func() {
		output, err := search.Search(10000, 5040)
		if err != nil {
			Fail("error should be nil")
		}
		Expect(output).To(Equal(int64(5040)))
	})

	It("should find best witness successfully", func() {
		count_till := int64(100_000)

		output, witnessVal := search.BestWitness(count_till, 11000)
		fmt.Println("\nCurrent Best till", humanize.Comma(int64(count_till)), "is", output, "at value", witnessVal)

		Expect(output).To(Equal(int64(55440)))
	})

	It("Should compute riemann sums correctly", func() {
		expectedOutput := []riemann.RiemannDivisorSum{
			{N: *big.NewInt(10080), DivisorSum: *big.NewInt(39312), WitnessValue: 1.75581},
			{N: *big.NewInt(10081), DivisorSum: *big.NewInt(10692), WitnessValue: 0.47749},
			{N: *big.NewInt(10082), DivisorSum: *big.NewInt(15339), WitnessValue: 0.68495},
		}

		actualOutput := []riemann.RiemannDivisorSum{
			search.NewExhaustiveSearchState(10080).ComputeRiemannDivisorSum(),
			search.NewExhaustiveSearchState(10081).ComputeRiemannDivisorSum(),
			search.NewExhaustiveSearchState(10082).ComputeRiemannDivisorSum(),
		}
		Expect(len(actualOutput)).To(Equal(len(expectedOutput)))
		for key, value := range actualOutput {
			Expect(value.DivisorSum).To(Equal(expectedOutput[key].DivisorSum))
			Expect(value.N).To(Equal(expectedOutput[key].N))
			Expect(math.Abs(value.WitnessValue-expectedOutput[key].WitnessValue) < 1e-5).To(BeTrue())
		}

	})

	// It("Should panic if it can't compute riemann sums", func() {
	// 	Expect(func() { riemann.ComputerRiemannDivisorSums(0, 1) }).Should(PanicWith("Divisor Sum cannot be found"))

	// })
})
