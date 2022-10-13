package search_test

import (
	"math/big"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/search"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Search Strategies", func() {

	Context("Exhaustive Search Strategies", func() {

		It("should find next states correctly across levels", func() {
			startingState := search.NewExhaustiveSearchState(10000)
			actualOutput := startingState.GetNextBatch(5)
			expectedOutput := []search.SearchState{
				search.NewExhaustiveSearchState(10001),
				search.NewExhaustiveSearchState(10002),
				search.NewExhaustiveSearchState(10003),
				search.NewExhaustiveSearchState(10004),
				search.NewExhaustiveSearchState(10005),
			}
			Expect(expectedOutput).To(Equal(actualOutput))

		})
	})

	Context("Superabundant Search Strategies", func() {

		It("should find next states correctly across levels", func() {
			startingState := search.NewSuperAbundantSearchState(2, 1, []int{})
			actualOutput := startingState.GetNextBatch(5)
			expectedOutput := []search.SearchState{
				search.NewSuperAbundantSearchState(3, 0, []int{3}),
				search.NewSuperAbundantSearchState(3, 1, []int{2, 1}),
				search.NewSuperAbundantSearchState(3, 2, []int{1, 1, 1}),
				search.NewSuperAbundantSearchState(4, 0, []int{4}),
				search.NewSuperAbundantSearchState(4, 1, []int{3, 1}),
			}
			Expect(expectedOutput).To(Equal(actualOutput))

		})

		It("should panic for illegal levels", func() {
			startingState := search.NewSuperAbundantSearchState(2, 2, []int{})
			Expect(func() { startingState.GetNextBatch(5) }).To(PanicWith("index level is illegal"))
		})
	})

	Context("Find N from prime factors", func() {

		It("should work correctly ", func() {
			inputArray := [][]int64{{2, 3}, {2, 2}}

			Expect(search.FindNFromPrimeFactors(inputArray)).To(Equal(*big.NewInt(32)))

		})
	})

	DescribeTable("Creates New Search State of both types", func(serializedState, stateType string, expectedOutput search.SearchState) {
		Expect(search.NewSearchState(serializedState, stateType)).To(Equal(expectedOutput))
	},
		Entry("Superabundant", "18, 161", "superabundant", search.SearchState(search.NewSuperAbundantSearchState(18, 161, []int{-1}))),
		Entry("Exhaustive", "10000", "exhaustive", search.SearchState(search.NewExhaustiveSearchState(10000))),
	)

})
