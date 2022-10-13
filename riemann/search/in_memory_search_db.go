package search

import (
	"fmt"
)

type InMemorySearchDb struct {
	data []SearchMetadata
}

func (imsdb *InMemorySearchDb) Initialize() {
	imsdb.data = []SearchMetadata{}
}

func (imsdb *InMemorySearchDb) LatestSearchState(stateType string) SearchState {

	for i := range imsdb.data {
		current := imsdb.data[len(imsdb.data)-i-1]
		if current.stateType == stateType {
			return current.endingState
		}
	}

	return InitialSearchState(stateType)

}

func (imsdb *InMemorySearchDb) InsertSearchMetadata(smd SearchMetadata) {
	imsdb.data = append(imsdb.data, smd)
}

func (imsdb *InMemorySearchDb) Close() {
	fmt.Println("Inmemory Search DB Closed!")
}
