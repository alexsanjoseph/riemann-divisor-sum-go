package riemann

import "time"

type SearchState interface {
	Serialize() string
	Value() int64
	GetNextBatch(int64) []SearchState
}

type SearchMetadata struct {
	startTime     time.Time
	endTime       time.Time
	stateType     string // convert to enum
	startingState SearchState
	endingState   SearchState
}

type SearchStateDB interface {
	Initialize()
	LatestSearchState(string) SearchState
	InsertSearchMetadata(SearchMetadata)
	Close()
}
