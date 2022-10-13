package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CounterExample Search", func() {

	Context("First N Primes", func() {
		It("should give empty list for 0", func() {
			Expect(riemann.FirstNPrimes(0)).To(Equal([]int{}))
		})

		It("should give 2 for 1", func() {
			Expect(riemann.FirstNPrimes(1)).To(Equal([]int{2}))
		})

		It("should give first 10 primes correctly", func() {
			expectedOutput := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
			Expect(riemann.FirstNPrimes(10)).To(Equal(expectedOutput))
		})

		It("should error out for values > 1000", func() {
			Expect(func() { riemann.FirstNPrimes(1001) }).To(PanicWith("value of primes too high"))
		})
	})
})
