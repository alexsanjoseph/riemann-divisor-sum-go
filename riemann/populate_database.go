package riemann

import (
	"fmt"
	"time"
)

func PopulateDB(db DivisorDb, sdb SearchStateDB, stateType string, batchSize int64, batches int64, minWitnessValue float64) {

	latestSearchState := sdb.LatestSearchState(stateType)
	nextBatch := latestSearchState.GetNextBatch(batchSize)

	for i := 0; batches == -1 || i < int(batches); i++ {
		time1 := time.Now()
		candidateResults := []RiemannDivisorSum{}

		for _, candidate := range nextBatch {
			candidateResult := candidate.ComputeRiemannDivisorSum()
			if candidateResult.WitnessValue > minWitnessValue {
				candidateResults = append(candidateResults, candidateResult)
			}
		}
		db.Upsert(candidateResults)
		endingState := nextBatch[len(nextBatch)-1]

		fmt.Printf("Computed Sums from %s to %s in %s \n",
			nextBatch[0].Serialize(), endingState.Serialize(), time.Since(time1))

		time2 := time.Now()

		singleBatchMetadata := SearchMetadata{
			startTime:     time1,
			endTime:       time2,
			stateType:     stateType,
			startingState: nextBatch[0],
			endingState:   endingState,
		}

		sdb.InsertSearchMetadata(singleBatchMetadata)
		fmt.Printf("Inserted Metadata in %s\n", time.Since(time2))

		time3 := time.Now()

		nextBatch = endingState.GetNextBatch(batchSize)
		fmt.Printf("Next batch computed in %s\n", time.Since(time3))

	}
}
