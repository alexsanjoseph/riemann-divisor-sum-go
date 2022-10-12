package riemann

import (
	"fmt"
)

type InMemorySearchDb struct {
	data []SearchMetadata
}

func (imsdb *InMemorySearchDb) Initialize() {
	imsdb.data = []SearchMetadata{}
}

func (imsdb *InMemorySearchDb) LatestSearchState(searchType string) SearchState {

	for i := range imsdb.data {
		current := imsdb.data[len(imsdb.data)-i-1]
		if current.stateType == searchType {
			return current.endingState
		}
	}

	if searchType == "exhaustive" {
		latestSearchState := SearchState(NewExhaustiveSearchState(10000))
		return latestSearchState
	}
	if searchType == "superabundant" {
		latestSearchState := SearchState(NewSuperAbundantSearchState(14, 0, []int{1}))
		return latestSearchState
	}

	panic("unknown searchType")
}

func (imsdb *InMemorySearchDb) InsertSearchMetadata(smd SearchMetadata) {
	imsdb.data = append(imsdb.data, smd)
}

func (imsdb *InMemorySearchDb) Close() {
	fmt.Println("Inmemory Search DB Closed!")
}
