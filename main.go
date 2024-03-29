package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/divisordb"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/populate"
	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann/search"
)

var (
	batchSize       = flag.Int64("batchSize", 1000000, "size per batch")
	stateType       = flag.String("stateType", "superabundant", "file-name")
	totalBatches    = flag.Int64("totalBatches", -1, "number of batches to calculate (-1) = Inf")
	minWitnessValue = flag.Float64("minWitnessValue", 1.75, "minimum Witness Value to persist")
)

func handleClose(db divisordb.DivisorDb) {
	fmt.Println("Got Interrupt. Shutting down...")
	summarizeOutput := db.Summarize()
	fmt.Println("\nHighest Number Analyzed\n======")

	fmt.Print(summarizeOutput.LargestComputedN.Print())
	fmt.Println("\nLargest Witness Value\n======")
	fmt.Print(summarizeOutput.LargestWitnessValue.Print())

}

func main() {

	flag.Parse()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	sqlDB := divisordb.SqliteDivisorDb{DBPath: "divisorDb.sqlite"}
	var ddb = divisordb.DivisorDb(&sqlDB)

	// imDB := riemann.InMemoryDivisorDb{}
	// var ddb = riemann.DivisorDb(&imDB)

	ddb.Initialize()
	defer ddb.Close()

	sqlsdb := search.SqliteSearchDb{DBPath: "searchDb.sqlite"}
	ssdb := search.SearchStateDB(&sqlsdb)

	// imsdb := riemann.InMemorySearchDb{}
	// ssdb := riemann.SearchStateDB(&imsdb)

	ssdb.Initialize()
	defer ssdb.Close()

	go populate.PopulateDB(ddb, ssdb, *stateType, *batchSize, *totalBatches, *minWitnessValue)
	<-sigCh
	handleClose(ddb)
}
