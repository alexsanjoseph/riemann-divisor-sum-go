package riemann

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

func WitnessValue(n int64, pds int64) float64 {
	denom := float64(n) * math.Log(math.Log(float64(n)))
	var divSum int64
	var err error
	if pds < 0 {
		divSum, err = DivisorSum(n)
	} else {
		divSum, err = pds, nil
	}
	if err != nil {
		panic(fmt.Sprintf("Error calculating DivisorSum for %d", n))
	}

	return float64(divSum) / float64(denom)
}

func Search(maxRange, searchStart int64) (int64, error) {
	for i := searchStart; i < maxRange; i++ {
		if WitnessValue(i, -1) > 1.782 {
			return i, nil
		}
	}
	return 0, errors.New("no witness value found")
}

// func BestWitness(maxRange, searchStart int64) (int64, float64) {
// 	maxVal := 0.0
// 	bestWitness := searchStart
// 	for i := searchStart; i < maxRange; i++ {
// 		currentWitness := WitnessValue(i, -1)
// 		if currentWitness > maxVal {
// 			bestWitness = i
// 			maxVal = currentWitness
// 		}
// 	}
// 	return bestWitness, maxVal
// }

type witness struct {
	idx   int64
	value float64
}

func BestWitness(maxRange, searchStart int64) (int64, float64) {
	parallelism := int64(100)
	allValues := make(chan witness, parallelism)
	var wg sync.WaitGroup
	wg.Add(int(parallelism))
	for i := int64(0); i < parallelism; i++ {
		go func(j int64) {
			defer wg.Done()
			maxVal := 0.0
			bestWitness := searchStart
			for k := searchStart + j; k < maxRange; k += parallelism {
				currentWitness := WitnessValue(k, -1)
				if currentWitness > maxVal {
					maxVal = currentWitness
					bestWitness = k
				}
			}
			allValues <- witness{bestWitness, maxVal}
		}(i)
	}
	wg.Wait()
	close(allValues)
	bestWitness := int64(0)
	maxVal := -1.0
	for val := range allValues {
		if val.value > maxVal {
			maxVal = val.value
			bestWitness = val.idx
		}
	}
	return bestWitness, maxVal
}

// func BestWitness(maxRange, searchStart int64) (int64, float64) {
// 	allValues := make(chan witness, maxRange-searchStart)
// 	var wg sync.WaitGroup
// 	wg.Add(int(maxRange - searchStart))
// 	for i := searchStart; i < maxRange; i++ {
// 		go func(j int64) {
// 			defer wg.Done()
// 			currentWitness := WitnessValue(j, -1)
// 			allValues <- witness{j, currentWitness}
// 		}(i)
// 	}
// 	wg.Wait()
// 	close(allValues)
// 	bestWitness := int64(0)
// 	maxVal := -1.0
// 	for val := range allValues {
// 		if val.value > maxVal {
// 			maxVal = val.value
// 			bestWitness = val.idx
// 		}
// 	}
// 	return bestWitness, maxVal
// }

func ComputerRiemannDivisorSums(startingN, endingN int64) []RiemannDivisorSum {
	output := []RiemannDivisorSum{}

	for i := startingN; i <= endingN; i++ {
		ds, err := DivisorSum(i)
		if err != nil {
			panic("Divisor Sum cannot be found")
		}
		wv := WitnessValue(i, ds)
		output = append(output, RiemannDivisorSum{N: i, WitnessValue: wv, DivisorSum: ds})
	}
	return output
}
