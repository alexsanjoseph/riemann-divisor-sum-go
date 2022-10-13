package riemann_test

import (
	"math/big"
	"os"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = AfterEach(func() {
	os.Remove(SearchDBPath)
	os.Remove(DivisorDBPath)
})

var _ = Describe("Parametrized Population tests", func() {

	DescribeTable("Populates, Summarizes and finds StartingN correctly", func(inputDb riemann.DivisorDb, searchDb riemann.SearchStateDB) {
		db := setupDivisorDB(inputDb)
		sdb := setupSearchStateDB(searchDb)

		riemann.PopulateDB(db, sdb, "exhaustive", 90, 1, 0)
		summaryData := db.Summarize()

		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(*big.NewInt(10080)))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(*big.NewInt(10090)))

		nextBatch := sdb.LatestSearchState("exhaustive").GetNextBatch(100)
		Expect(nextBatch[0].Value()).To(BeEquivalentTo("10091"))

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}, &riemann.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}, &riemann.InMemorySearchDb{}),
	)

	DescribeTable("Populates, Summarizes correctly for superabundant search", func(inputDb riemann.DivisorDb, searchDb riemann.SearchStateDB) {
		db := setupDivisorDB(inputDb)
		sdb := setupSearchStateDB(searchDb)

		riemann.PopulateDB(db, sdb, "superabundant", 90, 1, 0)
		summaryData := db.Summarize()

		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(*big.NewInt(10080)))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(*big.NewInt(6469693230)))

		nextBatch := sdb.LatestSearchState("superabundant").GetNextBatch(100)
		Expect(nextBatch[0].Serialize()).To(BeEquivalentTo("11, 19"))

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}, &riemann.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}, &riemann.InMemorySearchDb{}),
	)
})
