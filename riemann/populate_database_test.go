package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Populates Database correctly", func() {

	var db = riemann.DivisorDb(riemann.InMemoryDivisorDb{Data: make(map[int64]riemann.RiemannDivisorSum)})
	It("Populates and Summarizes correctly", func() {

		riemann.PopulateDB(db, 10070, 10085, 21)
		summaryData := db.Summarize()
		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(10080))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(10091))
	})

	It("Finds startingN correctly", func() {

		startingN := riemann.FindStartingNForDB(db, 10075)
		Expect(startingN).To(BeEquivalentTo(10092))

	})
})
