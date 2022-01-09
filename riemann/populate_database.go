package riemann

import "fmt"

func PopulateDB(db DivisorDb, startingN, endingN, batchSize int64) {
	var currentStartingN int64

	dbStartingN := db.Summarize().LargestComputedN.N

	if dbStartingN > startingN {
		currentStartingN = dbStartingN
	} else {
		currentStartingN = startingN
	}
	currentEndingN := currentStartingN + batchSize

	fmt.Println(dbStartingN, startingN, endingN, currentEndingN, currentStartingN)
	for endingN == -1 || currentEndingN < endingN+batchSize {
		db.Upsert(ComputerRiemannDivisorSums(currentStartingN, currentEndingN))
		fmt.Printf("Computed Sums from %d to %d\n", currentStartingN, currentEndingN)
		currentStartingN = currentEndingN + 1
		currentEndingN = currentStartingN + batchSize - 1
	}
}
