package riemann

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDivisorDb struct {
	DBPath string
	db     *sql.DB
}

func (sqdb *SqliteDivisorDb) Initialize() {

	db, err := sql.Open("sqlite3", sqdb.DBPath)
	sqdb.db = db
	if err != nil {
		panic(err)
	}

	sqlStmt := `
	CREATE TABLE RiemannDivisorSums (
            n TEXT CONSTRAINT divisor_sum_pk PRIMARY KEY,
            divisor_sum TEXT,
            witness_value REAL
	);
	`
	_, err = sqdb.db.Exec(sqlStmt)

	if err != nil && err.Error() != "table RiemannDivisorSums already exists" {
		panic(err)
	}
}

func (sqdb SqliteDivisorDb) Load() []RiemannDivisorSum {
	sqlStmt := `
            SELECT n, divisor_sum, witness_value
            FROM RiemannDivisorSums
            ORDER BY n asc;
			`
	rows, err := sqdb.db.Query(sqlStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	output := []RiemannDivisorSum{}
	for rows.Next() {
		var n, divisorSum string
		var witnessValue float64
		err = rows.Scan(&n, &divisorSum, &witnessValue)
		if err != nil {
			log.Fatal(err)
		}
		N, ok := new(big.Int).SetString(n, 10)
		if !ok {
			log.Fatal("unable to parse N")
		}

		DivisorSum, ok := new(big.Int).SetString(n, 10)
		if !ok {
			log.Fatal("unable to parse divisor sum")
		}

		output = append(output, RiemannDivisorSum{
			N:            *N,
			DivisorSum:   *DivisorSum,
			WitnessValue: witnessValue,
		})
	}
	return output
}

func GetStableTextRepresentationOfBigInt(N big.Int, fixedLength int) string {
	NString := N.String()
	paddingRequired := fixedLength - len(NString) // No Need to worry about weird characters
	if paddingRequired < 0 {
		panic("number is bigger than can be represented by string")
	}
	for i := 0; i < paddingRequired; i++ {
		NString = "0" + NString
	}
	return NString
}

func (sqdb SqliteDivisorDb) Upsert(rds []RiemannDivisorSum) {

	tx, err := sqdb.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt, err := tx.Prepare(`
		INSERT INTO
            RiemannDivisorSums(n, divisor_sum, witness_value)
            VALUES(?, ?, ?)
        ON CONFLICT(n) DO UPDATE SET
            divisor_sum=excluded.divisor_sum,
            witness_value=excluded.witness_value;
			;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()
	for _, value := range rds {
		_, err := sqlStmt.Exec(
			GetStableTextRepresentationOfBigInt(value.N, 100),
			GetStableTextRepresentationOfBigInt(value.DivisorSum, 100),
			fmt.Sprintf("%f", value.WitnessValue),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func (sqdb SqliteDivisorDb) Summarize() SummaryStats {

	largestNStms := `
	SELECT *
	FROM RiemannDivisorSums
	ORDER BY n DESC
	LIMIT 1;
	`
	row := sqdb.db.QueryRow(largestNStms)
	var n, divisorSum string
	var witnessValue float64
	err := row.Scan(&n, &divisorSum, &witnessValue)
	if err != nil {
		return SummaryStats{
			LargestWitnessValue: RiemannDivisorSum{},
			LargestComputedN:    RiemannDivisorSum{},
		}
	}

	N, ok := new(big.Int).SetString(n, 10)
	if !ok {
		log.Fatal("unable to parse N")
	}

	DivisorSum, ok := new(big.Int).SetString(divisorSum, 10)
	if !ok {
		log.Fatal("unable to parse divisor sum")
	}

	largest_computed_n := RiemannDivisorSum{
		N:            *N,
		DivisorSum:   *DivisorSum,
		WitnessValue: witnessValue,
	}

	largestWitnessStmt := `
	SELECT *
	FROM RiemannDivisorSums
	ORDER BY witness_value DESC
	LIMIT 1;
	`
	row = sqdb.db.QueryRow(largestWitnessStmt)
	err = row.Scan(&n, &divisorSum, &witnessValue)
	if err != nil {
		panic(err)
	}
	N, ok = new(big.Int).SetString(n, 10)
	if !ok {
		log.Fatal("unable to parse N")
	}

	DivisorSum, ok = new(big.Int).SetString(divisorSum, 10)
	if !ok {
		log.Fatal("unable to parse divisor sum")
	}

	largest_witness_value := RiemannDivisorSum{
		N:            *N,
		DivisorSum:   *DivisorSum,
		WitnessValue: witnessValue,
	}

	return SummaryStats{
		LargestWitnessValue: largest_witness_value,
		LargestComputedN:    largest_computed_n,
	}

}

func (sqdb *SqliteDivisorDb) Close() {
	err := sqdb.db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Closed!")
}
