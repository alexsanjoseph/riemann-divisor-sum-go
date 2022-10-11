package riemann

import (
	"fmt"
)

type InMemorySearchDb struct {
	data []SearchMetadata
}

func (imsdb *InMemorySearchDb) Initialize() {
	imsdb.data = []SearchMetadata{DefaultSearchMetadata()}
}

func (imsdb *InMemorySearchDb) LatestSearchState(searchType string) SearchState {
	return imsdb.data[len(imsdb.data)-1].endingState
}

func (imsdb *InMemorySearchDb) InsertSearchMetadata(smd SearchMetadata) {
	imsdb.data = append(imsdb.data, smd)
}

func (imsdb *InMemorySearchDb) Close() {
	fmt.Println("Inmemory Search DB Closed!")
}
