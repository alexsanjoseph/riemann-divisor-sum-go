package riemann

import "fmt"

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
		db.Upsert(ComputerRiemannDivisorSums(currentStartingN, currentEndingN))
		fmt.Printf("Computed Sums from %d to %d\n", currentStartingN, currentEndingN)
		currentStartingN = currentEndingN + 1
		currentEndingN = currentStartingN + batchSize - 1
	}
}
