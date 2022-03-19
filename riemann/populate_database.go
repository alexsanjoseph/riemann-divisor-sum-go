package riemann

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

func FindStartingNForDB(db DivisorDb, startingN int64) int64 {
	var currentStartingN int64

	dbStartingN := db.Summarize().LargestComputedN.N

	if dbStartingN > startingN {
		currentStartingN = dbStartingN + 1
	} else {
		currentStartingN = startingN
	}
	return currentStartingN
}

func PopulateDB(db DivisorDb, startingN, endingN, batchSize int64) {
	currentStartingN := FindStartingNForDB(db, startingN)
	currentEndingN := currentStartingN + batchSize

	for endingN == -1 || currentEndingN < endingN+batchSize {
		start := time.Now()
		db.Upsert(ComputerRiemannDivisorSums(currentStartingN, currentEndingN))
		currentStartingN = currentEndingN + 1
		currentEndingN = currentStartingN + batchSize - 1
		elapsed := time.Since(start)
		fmt.Printf("Computed Sums from %s to %s in %s \n",
			humanize.Comma(currentStartingN), humanize.Comma(currentEndingN), elapsed)
	}
}
