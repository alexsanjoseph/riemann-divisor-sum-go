package search

import (
	"os"
	"time"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
)

type SearchState interface {
	Serialize() string
	Value() string
	GetNextBatch(int64) []SearchState
	ComputeRiemannDivisorSum() riemann.RiemannDivisorSum
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

func (smm *SearchMetadata) Initialize(stateType string, startingState, endingState SearchState, startTime, endTime time.Time) {
	smm.startTime = startTime
	smm.endTime = endTime
	smm.stateType = stateType
	smm.startingState = startingState
	smm.endingState = endingState
}

func SetupSearchStateDB(inputDb SearchStateDB, SearchDBPath string) SearchStateDB {
	os.Remove(SearchDBPath)
	db := SearchStateDB(inputDb)
	db.Initialize()
	return db
}
