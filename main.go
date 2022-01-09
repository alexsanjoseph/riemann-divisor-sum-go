package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
)

func handleClose(db riemann.DivisorDb) {
	fmt.Println("Got Interrupt. Shutting down...")
	summarizeOutput := db.Summarize()
	fmt.Println("\nHighest Number Analyzed\n======")

	fmt.Printf("%+v\n", summarizeOutput.LargestComputedN)
	fmt.Println("\nLargest Witness Value\n======")
	fmt.Printf("%+v\n", summarizeOutput.LargestWitnessValue)

}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	var db = riemann.DivisorDb(riemann.InMemoryDivisorDb{Data: make(map[int]riemann.RiemannDivisorSum)})
	go riemann.PopulateDB(db, 1000000)

	<-sigCh
	handleClose(db)
}
