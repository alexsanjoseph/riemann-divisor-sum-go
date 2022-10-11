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

func DefaultSearchMetadata() SearchMetadata {
	return SearchMetadata{
		startTime:     time.Now(),
		endTime:       time.Now(),
		stateType:     "exhaustive", // Need to genericise this later
		startingState: SearchState(NewExhaustiveSearchState(1)),
		endingState:   SearchState(NewExhaustiveSearchState(10000)),
	}
}
