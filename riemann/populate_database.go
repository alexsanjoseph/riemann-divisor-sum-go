package riemann

import "fmt"

func PopulateDB(db DivisorDb, batchSize int64) {
	var startingN int64
	dbStartingN := db.Summarize().LargestComputedN.N
	if dbStartingN < 5040 {
		startingN = 5041
	} else {
		startingN = dbStartingN
	}

	for {
		endingN := startingN + batchSize - 1
		db.Upsert(ComputerRiemannDivisorSums(startingN, endingN))
		fmt.Printf("Computed Sums from %d to %d\n", startingN, endingN)
		startingN = endingN + 1
	}
}
