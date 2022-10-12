package riemann_test

import (
	"os"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const SearchDBPath = "testSearchDB.sqlite"

func setupSearchStateDB(inputDb riemann.SearchStateDB) riemann.SearchStateDB {
	os.Remove(SearchDBPath)
	db := riemann.SearchStateDB(inputDb)
	db.Initialize()
	return db
}

var _ = AfterEach(func() {
	os.Remove(SearchDBPath)
})

var _ = BeforeEach(func() {

})

var _ = Describe("Parametrized Search Database Tests", func() {

	DescribeTable("is initially creates bootstrap data", func(searchDb riemann.SearchStateDB) {

		db := setupSearchStateDB(searchDb)
		loadedData := db.LatestSearchState("superabundant")
		actualOutput := loadedData.Value()
		Expect(actualOutput).To(Equal("[1]"))

	},
		Entry("SQLite", &riemann.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &riemann.InMemorySearchDb{}),
	)

	DescribeTable("inserts correctly", func(searchDb riemann.SearchStateDB) {
		db := setupSearchStateDB(searchDb)
		newSearchMetadata := riemann.SearchMetadata{}
		startingState := riemann.NewSearchState("14, 1", "superabundant")
		endingState := riemann.NewSearchState("16, 1", "superabundant")
		newSearchMetadata.Initialize("superabundant", startingState, endingState)
		db.InsertSearchMetadata(newSearchMetadata)
		actualOutput := db.LatestSearchState("superabundant")
		Expect(actualOutput).To(Equal(riemann.NewSuperAbundantSearchState(16, int64(1), []int{-1})))

	},
		Entry("SQLite", &riemann.SqliteSearchDb{DBPath: SearchDBPath}),
		Entry("In-Memory", &riemann.InMemorySearchDb{}),
	)

	// DescribeTable("get latest search state", func(searchDb riemann.SearchStateDB) {

	// },
	// 	Entry("SQLite", &riemann.SqliteSearchDb{DBPath: SearchDBPath}),
	// 	Entry("In-Memory", &riemann.InMemorySearchDb{}),
	// )

})
