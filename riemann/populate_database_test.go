package riemann_test

import (
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parametrized Population tests", func() {

	DescribeTable("Populates, Summarizes and finds StartingN correctly", func(inputDb riemann.DivisorDb, searchDb riemann.SearchStateDB) {
		db := setupDivisorDB(inputDb)
		sdb := setupSearchStateDB(searchDb)

		riemann.PopulateDB(db, sdb, 90, 1)
		summaryData := db.Summarize()

		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(10080))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(10090))

		nextBatch := sdb.LatestSearchState("exhaustive").GetNextBatch(100)
		Expect(nextBatch[0].Value()).To(BeEquivalentTo(10091))

	},
		// Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}, &riemann.InMemorySearchDb{}),
	)
})
