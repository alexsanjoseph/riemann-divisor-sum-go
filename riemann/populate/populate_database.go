package populate

import (
	"fmt"
	"time"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/divisordb"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/search"
)

func PopulateDB(db divisordb.DivisorDb, sdb search.SearchStateDB, stateType string, batchSize int64, batches int64, minWitnessValue float64) {

	latestSearchState := sdb.LatestSearchState(stateType)
	nextBatch := latestSearchState.GetNextBatch(batchSize)

	for i := 0; batches == -1 || i < int(batches); i++ {
		time1 := time.Now()
		candidateResults := []riemann.RiemannDivisorSum{}

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

		singleBatchMetadata := search.SearchMetadata{}
		singleBatchMetadata.Initialize(
			stateType,
			nextBatch[0],
			endingState,
			time1,
			time2,
		)

		sdb.InsertSearchMetadata(singleBatchMetadata)
		fmt.Printf("Inserted Metadata in %s\n", time.Since(time2))

		time3 := time.Now()

		nextBatch = endingState.GetNextBatch(batchSize)
		fmt.Printf("Next batch computed in %s\n", time.Since(time3))

	}
}
