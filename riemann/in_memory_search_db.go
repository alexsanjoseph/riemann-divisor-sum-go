package riemann

import (
	"fmt"
	"time"
)

type InMemorySearchDb struct {
	data []SearchMetadata
}

func (imsdb *InMemorySearchDb) Initialize() {
	imsdb.data = []SearchMetadata{SearchMetadata{
		startTime:     time.Now(),
		endTime:       time.Now(),
		stateType:     "exhaustive",
		startingState: NewExhaustiveSearchState(1),
		endingState:   NewExhaustiveSearchState(10000),
	}}
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
