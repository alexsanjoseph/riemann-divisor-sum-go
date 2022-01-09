package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Divisor", func() {

	Describe("Good Cases", func() {
		It("Should give correct result", func() {
			output, err := riemann.DivisorSum(72)
			if err != nil {
				Fail("error should be nil")
			}
			Expect(output).To(Equal(int64(195)))
		})

		It("Should give correct result", func() {
			output, err := riemann.DivisorSum(1)
			if err != nil {
				Fail("error should be nil")
			}
			Expect(output).To(Equal(int64(1)))
		})
	})

	Context("Bad Cases", func() {
		It("zero value", func() {
			output, err := riemann.DivisorSum(0)
			Expect(output).To(BeZero())
			Expect(err).To(HaveOccurred())
		})

		It("negative value", func() {
			output, err := riemann.DivisorSum(-5)
			Expect(output).To(BeZero())
			Expect(err).To(HaveOccurred())
		})
	})

	It("Should work on parametrized cases", func() {
		divisorSums := []int64{
			1, 3, 4, 7, 6, 12, 8, 15, 13, 18, 12, 28, 14, 24, 24, 31, 18, 39, 20, 42, 32,
			36, 24, 60, 31, 42, 40, 56, 30, 72, 32, 63, 48, 54, 48, 91, 38, 60, 56, 90, 42,
			96, 44, 84, 78, 72, 48, 124, 57, 93, 72,
		}

		for i, sum := range divisorSums {
			By("Passing each cases", func() {
				output, err := riemann.DivisorSum(int64(i) + 1)
				if err != nil {
					Fail("error should be nil")
				}
				Expect(output).To(Equal(sum))
			})
		}

	})
})
