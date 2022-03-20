package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parametrized Population tests", func() {

	DescribeTable("Populates, Summarizes and finds StartingN correctly", func(inputDb riemann.DivisorDb) {
		db := beforeEachFunc(inputDb)
		riemann.PopulateDB(db, 10070, 10085, 21)
		summaryData := db.Summarize()

		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(10080))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(10091))

		startingN := riemann.FindStartingNForDB(db, 10075)
		Expect(startingN).To(BeEquivalentTo(10092))

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)
})
