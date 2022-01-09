package riemann

type RiemannDivisorSum struct {
	N            int
	DivisorSum   int
	WitnessValue float64
}

type SummaryStats struct {
	LargestComputedN    RiemannDivisorSum
	LargestWitnessValue RiemannDivisorSum
}

type DivisorDb interface {
	Load() []RiemannDivisorSum
	Upsert([]RiemannDivisorSum)
	Summarize() SummaryStats
}
