package riemann_test

import (
	"math/big"
	"os"
	"sort"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const DivisorDBPath = "testDivisorDB.sqlite"

func setupDivisorDB(inputDb riemann.DivisorDb) riemann.DivisorDb {
	os.Remove(DivisorDBPath)
	db := riemann.DivisorDb(inputDb)
	db.Initialize()
	return db
}

var _ = AfterEach(func() {
	os.Remove(DivisorDBPath)
})

var _ = Describe("Parametrized Database Tests", func() {

	DescribeTable("is initially empty", func(inputDb riemann.DivisorDb) {
		db := setupDivisorDB(inputDb)
		loadedData := db.Load()
		Expect(len(loadedData)).To(Equal(0))

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Multiple initializations should be Idempotent", func(inputDb riemann.DivisorDb) {
		db := setupDivisorDB(inputDb)
		db.Initialize()
		db.Initialize()
		Expect(true)
	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Upserts correctly", func(inputDb riemann.DivisorDb) {
		db := setupDivisorDB(inputDb)
		records := []riemann.RiemannDivisorSum{
			{N: *big.NewInt(1), DivisorSum: *big.NewInt(1), WitnessValue: 1},
			{N: *big.NewInt(2), DivisorSum: *big.NewInt(2), WitnessValue: 2},
		}

		By("upserting fine from empty", func() {
			db.Upsert(records)
			loadedData := db.Load()
			sort.Slice(loadedData, func(p, q int) bool {
				return loadedData[p].N.Cmp(&loadedData[q].N) == -1
			})
			Expect(loadedData).To(Equal(records))
		})

		By("upserting fine from non-empty", func() {
			newRecords := []riemann.RiemannDivisorSum{
				{N: *big.NewInt(3), DivisorSum: *big.NewInt(3), WitnessValue: 3},
				{N: *big.NewInt(4), DivisorSum: *big.NewInt(4), WitnessValue: 4},
			}
			db.Upsert(newRecords)
			loadedData := db.Load()
			sort.Slice(loadedData, func(p, q int) bool {
				return loadedData[p].N.Cmp(&loadedData[q].N) == -1
			})
			Expect(loadedData).To(Equal(append(records, newRecords...)))
		})

		By("overriding existing docs when upserted", func() {
			newRecords := []riemann.RiemannDivisorSum{
				{N: *big.NewInt(3), DivisorSum: *big.NewInt(3), WitnessValue: 10},
				{N: *big.NewInt(5), DivisorSum: *big.NewInt(5), WitnessValue: 5},
			}
			expectedNewRecords := []riemann.RiemannDivisorSum{
				{N: *big.NewInt(3), DivisorSum: *big.NewInt(3), WitnessValue: 10},
				{N: *big.NewInt(4), DivisorSum: *big.NewInt(4), WitnessValue: 4},
				{N: *big.NewInt(5), DivisorSum: *big.NewInt(5), WitnessValue: 5},
			}
			db.Upsert(newRecords)
			loadedData := db.Load()
			sort.Slice(loadedData, func(p, q int) bool {
				return loadedData[p].N.Cmp(&loadedData[q].N) == -1
			})
			Expect(loadedData).To(Equal(append(records, expectedNewRecords...)))
		})

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Summarizes", func(inputDb riemann.DivisorDb) {
		db := setupDivisorDB(inputDb)
		By("correctly summarizing empty data", func() {
			summaryData := db.Summarize()
			expectedSummaryData := riemann.SummaryStats{
				LargestWitnessValue: riemann.RiemannDivisorSum{},
				LargestComputedN:    riemann.RiemannDivisorSum{},
			}
			Expect(summaryData).To(Equal(expectedSummaryData))
		})

		By("correctly summarizing non-empty data", func() {
			records := []riemann.RiemannDivisorSum{
				{N: *big.NewInt(1), DivisorSum: *big.NewInt(1), WitnessValue: 10},
				{N: *big.NewInt(2), DivisorSum: *big.NewInt(2), WitnessValue: 20},
				{N: *big.NewInt(3), DivisorSum: *big.NewInt(2), WitnessValue: 3},
			}
			db.Upsert(records)
			summaryData := db.Summarize()
			expectedSummaryData := riemann.SummaryStats{
				LargestWitnessValue: riemann.RiemannDivisorSum{N: *big.NewInt(2), DivisorSum: *big.NewInt(2), WitnessValue: 20},
				LargestComputedN:    riemann.RiemannDivisorSum{N: *big.NewInt(3), DivisorSum: *big.NewInt(2), WitnessValue: 3},
			}
			Expect(summaryData).To(Equal(expectedSummaryData))
		})

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Summarizes for float values", func(inputDb riemann.DivisorDb) {
		db := setupDivisorDB(inputDb)
		By("correctly summarizing non-empty data", func() {
			records := []riemann.RiemannDivisorSum{
				{N: *big.NewInt(10092), DivisorSum: *big.NewInt(24388), WitnessValue: 1.088},
				{N: *big.NewInt(10080), DivisorSum: *big.NewInt(39000), WitnessValue: 1.788},
			}
			db.Upsert(records)
			summaryData := db.Summarize()
			expectedSummaryData := riemann.SummaryStats{
				LargestWitnessValue: riemann.RiemannDivisorSum{N: *big.NewInt(10080), DivisorSum: *big.NewInt(39000), WitnessValue: 1.788},
				LargestComputedN:    riemann.RiemannDivisorSum{N: *big.NewInt(10092), DivisorSum: *big.NewInt(24388), WitnessValue: 1.088},
			}
			Expect(summaryData).To(Equal(expectedSummaryData))
		})
	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DivisorDBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	Describe("Represents Bigint Correctly in DB", func() {
		It("should work correctly", func() {
			actualOutput := riemann.GetStableTextRepresentationOfBigInt(*big.NewInt(1), 10)
			expectedOutput := "0000000001"
			Expect(actualOutput).To(Equal(expectedOutput))
		})

		It("should fail in large cases correctly", func() {
			Expect(func() { riemann.GetStableTextRepresentationOfBigInt(*big.NewInt(100), 2) }).To(PanicWith("number is bigger than can be represented by string"))
		})

	})

})
