package riemann

import "time"

type SearchState interface {
	Serialize() string
	Value() string
	GetNextBatch(int64) []SearchState
	ComputeRiemannDivisorSum() RiemannDivisorSum
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

func (smm *SearchMetadata) Initialize(stateType string, startingState, endingState SearchState) {
	smm.startTime = time.Now()
	smm.endTime = time.Now()
	smm.stateType = stateType
	smm.startingState = startingState
	smm.endingState = endingState
}
