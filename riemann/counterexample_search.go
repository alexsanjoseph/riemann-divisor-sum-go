package riemann

import (
	"errors"
	"math"
)

func WitnessValue(n int, pds int) float64 {
	denom := float64(n) * math.Log(math.Log(float64(n)))
	var divSum int
	var err error
	if pds < 0 {
		divSum, err = DivisorSum(n)
	} else {
		divSum, err = pds, nil
	}
	if err != nil {
		panic("Error calculating DivisorSum")
	}

	return float64(divSum) / float64(denom)
}

func Search(maxRange, searchStart int) (int, error) {
	for i := searchStart; i < maxRange; i++ {
		if WitnessValue(i, -1) > 1.782 {
			return i, nil
		}
	}
	return 0, errors.New("no witness value found")
}

func BestWitness(maxRange, searchStart int) (int, float64) {
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
