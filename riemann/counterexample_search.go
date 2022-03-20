package riemann

import (
	"errors"
	"fmt"
	"math"
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

func BestWitness(maxRange, searchStart int64) (int64, float64) {
	maxVal := 0.0
	bestWitness := searchStart
	for i := searchStart; i < maxRange; i++ {
		currentWitness := WitnessValue(i, -1)
		if currentWitness > maxVal {
			bestWitness = i
			maxVal = currentWitness
		}
	}
	return bestWitness, maxVal
}

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
