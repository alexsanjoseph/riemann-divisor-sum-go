package riemann

import "fmt"

type InMemoryDivisorDb struct {
	data map[int64]RiemannDivisorSum
}

func (imdb *InMemoryDivisorDb) Initialize() {
	imdb.data = map[int64]RiemannDivisorSum{}
}

func (imdb InMemoryDivisorDb) Load() []RiemannDivisorSum {
	output := []RiemannDivisorSum{}
	for _, value := range imdb.data {
		output = append(output, value)
	}
	return output
}

func (imdb InMemoryDivisorDb) Upsert(rds []RiemannDivisorSum) {
	for _, value := range rds {
		imdb.data[value.N] = value
	}
}

func (imdb InMemoryDivisorDb) Summarize() SummaryStats {
	if len(imdb.data) == 0 {
		return SummaryStats{
			LargestWitnessValue: RiemannDivisorSum{},
			LargestComputedN:    RiemannDivisorSum{},
		}
	}
	largest_computed_n := RiemannDivisorSum{N: -1}
	largest_witness_value := RiemannDivisorSum{WitnessValue: -1}

	for _, rds := range imdb.data {
		if rds.N > largest_computed_n.N {
			largest_computed_n = rds
		}
		if rds.WitnessValue > largest_witness_value.WitnessValue {
			largest_witness_value = rds
		}
	}
	return SummaryStats{
		LargestWitnessValue: largest_witness_value,
		LargestComputedN:    largest_computed_n,
	}

}

func (imdb *InMemoryDivisorDb) Close() {
	fmt.Println("Inmemory Divisor DB Closed!")
}
