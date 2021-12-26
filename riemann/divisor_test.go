package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Divisor", func() {

	Context("Good Cases", func() {
		It("Should give correct result", func() {
			output, _ := riemann.DivisorSum(72)
			Expect(output).To(Equal(195))
		})

		It("Should give correct result", func() {
			output, _ := riemann.DivisorSum(1)
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
