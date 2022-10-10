package riemann

import (
	"fmt"
	"time"
)

func PopulateDB(db DivisorDb, sdb SearchStateDB, batchSize int64, batches int64) {
	stateType := "exhaustive" //hardcoded for now

	latestSearchState := sdb.LatestSearchState(stateType)
	nextBatch := latestSearchState.GetNextBatch(batchSize)

	for i := 0; batches == -1 || i < int(batches); i++ {
		startTime := time.Now()
		candidateResults := []RiemannDivisorSum{}

		for _, candidate := range nextBatch {
			candidateResult := ComputeRiemannDivisorSum(candidate)
			candidateResults = append(candidateResults, candidateResult)
		}
		db.Upsert(candidateResults)

		endTime := time.Now()
		endingState := nextBatch[len(nextBatch)-1]

		elapsed := time.Since(startTime)
		fmt.Printf("Computed Sums from %s to %s in %s \n",
			nextBatch[0].Serialize(), endingState.Serialize(), elapsed)

		singleBatchMetadata := SearchMetadata{
			startTime:     startTime,
			endTime:       endTime,
			stateType:     stateType,
			startingState: nextBatch[0],
			endingState:   endingState,
		}
		sdb.InsertSearchMetadata(singleBatchMetadata)
		nextBatch = endingState.GetNextBatch(batchSize)
	}
}
