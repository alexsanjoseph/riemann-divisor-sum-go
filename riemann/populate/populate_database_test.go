package populate_test

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/divisordb"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/populate"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/search"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/types"
	. "github.com/onsi/gomega"
)

func TestDivisor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Populate Tests", types.ReporterConfig{
		SlowSpecThreshold: 100 * time.Millisecond,
	})
}

const DivisorDBPath = "testDivisorDB.sqlite"
const SearchDBPath = "testSearchDB.sqlite"

var _ = AfterEach(func() {
	os.Remove(SearchDBPath)
	os.Remove(DivisorDBPath)
})

var _ = Describe("Parametrized Population tests", func() {

	DescribeTable("Populates, Summarizes and finds StartingN correctly", func(inputDb divisordb.DivisorDb, searchDb search.SearchStateDB) {
		db := divisordb.SetupDivisorDB(inputDb, DivisorDBPath)
		sdb := search.SetupSearchStateDB(searchDb, SearchDBPath)

		populate.PopulateDB(db, sdb, "exhaustive", 90, 1, 0)
		summaryData := db.Summarize()

		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(*big.NewInt(10080)))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(*big.NewInt(10090)))

		nextBatch := sdb.LatestSearchState("exhaustive").GetNextBatch(100)
		Expect(nextBatch[0].Value()).To(BeEquivalentTo("10091"))

	},
		Entry("SQLite", &divisordb.SqliteDivisorDb{DBPath: DivisorDBPath}, &search.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &divisordb.InMemoryDivisorDb{}, &search.InMemorySearchDb{}),
	)

	DescribeTable("Populates, Summarizes correctly for superabundant search", func(inputDb divisordb.DivisorDb, searchDb search.SearchStateDB) {
		db := divisordb.SetupDivisorDB(inputDb, DivisorDBPath)
		sdb := search.SetupSearchStateDB(searchDb, SearchDBPath)

		populate.PopulateDB(db, sdb, "superabundant", 90, 1, 0)
		summaryData := db.Summarize()

		Expect(summaryData.LargestWitnessValue.N).To(BeEquivalentTo(*big.NewInt(10080)))
		Expect(summaryData.LargestComputedN.N).To(BeEquivalentTo(*big.NewInt(6469693230)))

		nextBatch := sdb.LatestSearchState("superabundant").GetNextBatch(100)
		Expect(nextBatch[0].Serialize()).To(BeEquivalentTo("11, 19"))

	},
		Entry("SQLite", &divisordb.SqliteDivisorDb{DBPath: DivisorDBPath}, &search.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &divisordb.InMemoryDivisorDb{}, &search.InMemorySearchDb{}),
	)
})
