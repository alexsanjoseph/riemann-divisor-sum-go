package search_test

import (
	"os"
	"time"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/search"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const SearchDBPath = "testSearchDB.sqlite"

var _ = AfterEach(func() {
	os.Remove(SearchDBPath)
})

var _ = BeforeEach(func() {

})

var _ = Describe("Parametrized Search Database Tests", func() {

	DescribeTable("is initially creates bootstrap data", func(searchDb search.SearchStateDB) {

		db := search.SetupSearchStateDB(searchDb, SearchDBPath)
		loadedData := db.LatestSearchState("superabundant")
		actualOutput := loadedData.Value()
		Expect(actualOutput).To(Equal("[1]"))

	},
		Entry("SQLite", &search.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &search.InMemorySearchDb{}),
	)

	DescribeTable("inserts correctly", func(searchDb search.SearchStateDB) {
		db := search.SetupSearchStateDB(searchDb, SearchDBPath)
		newSearchMetadata := search.SearchMetadata{}
		startingState := search.NewSearchState("14, 1", "superabundant")
		endingState := search.NewSearchState("16, 1", "superabundant")
		newSearchMetadata.Initialize("superabundant", startingState, endingState, time.Now(), time.Now())
		db.InsertSearchMetadata(newSearchMetadata)
		actualOutput := db.LatestSearchState("superabundant")
		Expect(actualOutput).To(Equal(search.NewSuperAbundantSearchState(16, int64(1), []int{-1})))

	},
		Entry("SQLite", &search.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &search.InMemorySearchDb{}),
	)

	// DescribeTable("get latest search state", func(searchDb riemann.SearchStateDB) {

	// },
	// 	Entry("SQLite", &riemann.SqliteSearchDb{DBPath: SearchDBPath}),
	// 	Entry("In-Memory", &riemann.InMemorySearchDb{}),
	// )

})
