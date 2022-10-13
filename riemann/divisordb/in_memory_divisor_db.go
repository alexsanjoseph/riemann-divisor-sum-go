package divisordb

import (
	"fmt"
	"math/big"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
)

type InMemoryDivisorDb struct {
	data map[string]riemann.RiemannDivisorSum
}

func (imdb *InMemoryDivisorDb) Initialize() {
	imdb.data = map[string]riemann.RiemannDivisorSum{}
}

func (imdb InMemoryDivisorDb) Load() []riemann.RiemannDivisorSum {
	output := []riemann.RiemannDivisorSum{}
	for _, value := range imdb.data {
		output = append(output, value)
	}
	return output
}

func (imdb InMemoryDivisorDb) Upsert(rds []riemann.RiemannDivisorSum) {
	for _, value := range rds {
		imdb.data[value.N.String()] = value
	}
}

func (imdb InMemoryDivisorDb) Summarize() riemann.SummaryStats {
	if len(imdb.data) == 0 {
		return riemann.SummaryStats{
			LargestWitnessValue: riemann.RiemannDivisorSum{},
			LargestComputedN:    riemann.RiemannDivisorSum{},
		}
	}
	largest_computed_n := riemann.RiemannDivisorSum{N: *big.NewInt(-1)}
	largest_witness_value := riemann.RiemannDivisorSum{WitnessValue: -1}

	for _, rds := range imdb.data {
		if rds.N.Cmp(&largest_computed_n.N) == 1 {
			largest_computed_n = rds
		}
		if rds.WitnessValue > largest_witness_value.WitnessValue {
			largest_witness_value = rds
		}
	}
	return riemann.SummaryStats{
		LargestWitnessValue: largest_witness_value,
		LargestComputedN:    largest_computed_n,
	}

}

func (imdb *InMemoryDivisorDb) Close() {
	fmt.Println("Inmemory Divisor DB Closed!")
}
