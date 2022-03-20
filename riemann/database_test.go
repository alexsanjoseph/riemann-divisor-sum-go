package riemann_test

import (
	"os"
	"sort"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const DBPath = "test.sqlite"

func beforeEachFunc(inputDb riemann.DivisorDb) riemann.DivisorDb {
	os.Remove(DBPath)
	db := riemann.DivisorDb(inputDb)
	db.Initialize()
	return db
}

var _ = AfterEach(func() {
	os.Remove(DBPath)
})

var _ = Describe("Parametrized Database Tests", func() {

	DescribeTable("is initially empty", func(inputDb riemann.DivisorDb) {
		db := beforeEachFunc(inputDb)
		loadedData := db.Load()
		Expect(len(loadedData)).To(Equal(0))

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Multiple initializations should be Idempotent", func(inputDb riemann.DivisorDb) {
		db := beforeEachFunc(inputDb)
		db.Initialize()
		db.Initialize()
		Expect(true)
	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Upserts correctly", func(inputDb riemann.DivisorDb) {
		db := beforeEachFunc(inputDb)
		records := []riemann.RiemannDivisorSum{
			{N: 1, DivisorSum: 1, WitnessValue: 1},
			{N: 2, DivisorSum: 2, WitnessValue: 2},
		}

		By("upserting fine from empty", func() {
			db.Upsert(records)
			loadedData := db.Load()
			sort.Slice(loadedData, func(p, q int) bool {
				return loadedData[p].N < loadedData[q].N
			})
			Expect(loadedData).To(Equal(records))
		})

		By("upserting fine from non-empty", func() {
			newRecords := []riemann.RiemannDivisorSum{
				{N: 3, DivisorSum: 3, WitnessValue: 3},
				{N: 4, DivisorSum: 4, WitnessValue: 4},
			}
			db.Upsert(newRecords)
			loadedData := db.Load()
			sort.Slice(loadedData, func(p, q int) bool {
				return loadedData[p].N < loadedData[q].N
			})
			Expect(loadedData).To(Equal(append(records, newRecords...)))
		})

		By("overriding existing docs when upserted", func() {
			newRecords := []riemann.RiemannDivisorSum{
				{N: 3, DivisorSum: 3, WitnessValue: 10},
				{N: 5, DivisorSum: 5, WitnessValue: 5},
			}
			expectedNewRecords := []riemann.RiemannDivisorSum{
				{N: 3, DivisorSum: 3, WitnessValue: 10},
				{N: 4, DivisorSum: 4, WitnessValue: 4},
				{N: 5, DivisorSum: 5, WitnessValue: 5},
			}
			db.Upsert(newRecords)
			loadedData := db.Load()
			sort.Slice(loadedData, func(p, q int) bool {
				return loadedData[p].N < loadedData[q].N
			})
			Expect(loadedData).To(Equal(append(records, expectedNewRecords...)))
		})

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Summarizes", func(inputDb riemann.DivisorDb) {
		db := beforeEachFunc(inputDb)
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
				{N: 1, DivisorSum: 1, WitnessValue: 10},
				{N: 2, DivisorSum: 2, WitnessValue: 20},
				{N: 3, DivisorSum: 2, WitnessValue: 3},
			}
			db.Upsert(records)
			summaryData := db.Summarize()
			expectedSummaryData := riemann.SummaryStats{
				LargestWitnessValue: riemann.RiemannDivisorSum{N: 2, DivisorSum: 2, WitnessValue: 20},
				LargestComputedN:    riemann.RiemannDivisorSum{N: 3, DivisorSum: 2, WitnessValue: 3},
			}
			Expect(summaryData).To(Equal(expectedSummaryData))
		})

	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

	DescribeTable("Summarizes for float values", func(inputDb riemann.DivisorDb) {
		db := beforeEachFunc(inputDb)
		By("correctly summarizing non-empty data", func() {
			records := []riemann.RiemannDivisorSum{
				{N: 10092, DivisorSum: 24388, WitnessValue: 1.088},
				{N: 10080, DivisorSum: 39000, WitnessValue: 1.788},
			}
			db.Upsert(records)
			summaryData := db.Summarize()
			expectedSummaryData := riemann.SummaryStats{
				LargestWitnessValue: riemann.RiemannDivisorSum{N: 10080, DivisorSum: 39000, WitnessValue: 1.788},
				LargestComputedN:    riemann.RiemannDivisorSum{N: 10092, DivisorSum: 24388, WitnessValue: 1.088},
			}
			Expect(summaryData).To(Equal(expectedSummaryData))
		})
	},
		Entry("SQLite", &riemann.SqliteDivisorDb{DBPath: DBPath}),
		Entry("In-Memory", &riemann.InMemoryDivisorDb{}),
	)

})
