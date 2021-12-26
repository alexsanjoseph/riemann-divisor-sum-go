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
			Expect(output).To(Equal(195))
		})

		It("Should give correct result", func() {
			output, err := riemann.DivisorSum(1)
			if err != nil {
				Fail("error should be nil")
			}
			Expect(output).To(Equal(1))
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
})
