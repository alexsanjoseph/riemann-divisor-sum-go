package riemann

import (
	"database/sql"
	"fmt"
	"log"

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
            n UNSIGNED BIG INT CONSTRAINT divisor_sum_pk PRIMARY KEY,
            divisor_sum UNSIGNED BIG INT,
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
		var n, divisorSum int64
		var witnessValue float64
		err = rows.Scan(&n, &divisorSum, &witnessValue)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, RiemannDivisorSum{
			N:            n,
			DivisorSum:   divisorSum,
			WitnessValue: witnessValue,
		})
	}
	return output
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
			fmt.Sprintf("%d", value.N),
			fmt.Sprintf("%d", value.DivisorSum),
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
	var n, divisorSum int64
	var witnessValue float64
	err := row.Scan(&n, &divisorSum, &witnessValue)
	if err != nil {
		return SummaryStats{
			LargestWitnessValue: RiemannDivisorSum{},
			LargestComputedN:    RiemannDivisorSum{},
		}
	}

	largest_computed_n := RiemannDivisorSum{
		N:            n,
		DivisorSum:   divisorSum,
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
	largest_witness_value := RiemannDivisorSum{
		N:            n,
		DivisorSum:   divisorSum,
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
