package riemann

import (
	"fmt"
	"math/big"
)

type RiemannDivisorSum struct {
	N            big.Int
	DivisorSum   big.Int
	WitnessValue float64
}

type SummaryStats struct {
	LargestComputedN    RiemannDivisorSum
	LargestWitnessValue RiemannDivisorSum
}

func (rds *RiemannDivisorSum) Print() string {
	return fmt.Sprintf(
		"Number: %s, DivisorSum: %s, WitnessValue %f\n",
		rds.N.String(), rds.DivisorSum.String(), rds.WitnessValue,
	)
}
