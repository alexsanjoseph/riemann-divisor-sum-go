package divisordb

import (
	"os"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
)

type DivisorDb interface {
	Initialize()
	Load() []riemann.RiemannDivisorSum
	Upsert([]riemann.RiemannDivisorSum)
	Summarize() riemann.SummaryStats
	Close()
}

func SetupDivisorDB(inputDb DivisorDb, DivisorDBPath string) DivisorDb {
	os.Remove(DivisorDBPath)
	db := DivisorDb(inputDb)
	db.Initialize()
	return db
}
