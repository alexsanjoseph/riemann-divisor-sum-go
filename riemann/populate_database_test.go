package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Populates Database correctly", func() {

	It("Populates and Summarizes correctly", func() {
		var db = riemann.DivisorDb(riemann.InMemoryDivisorDb{Data: make(map[int64]riemann.RiemannDivisorSum)})
		riemann.PopulateDB(db, 10070, 10085, 21)
		summaryData := db.Summarize()
		Expect(summaryData.LargestWitnessValue.N).To(Equal(int64(10080)))
	})
})
