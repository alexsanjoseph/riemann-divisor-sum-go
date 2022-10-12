package riemann

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteSearchDb struct {
	DBPath string
	db     *sql.DB
}

func (ssdb *SqliteSearchDb) Initialize() {
	db, err := sql.Open("sqlite3", ssdb.DBPath)
	ssdb.db = db
	if err != nil {
		panic(err)
	}

	sqlStmt := `
	CREATE TABLE SearchState (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			start_time TEXT,
			end_time TEXT,
			state_type TEXT,
			starting_state TEXT,
			ending_state TEXT
	);
	`
	_, err = ssdb.db.Exec(sqlStmt)

	if err != nil {
		if err.Error() == "table SearchState already exists" {
			return
		}
		panic(err)
	}
}

func (ssdb *SqliteSearchDb) LatestSearchState(searchType string) SearchState {
	sqlStmt := `
            SELECT ending_state
            FROM SearchState
            ORDER BY end_time DESC, id DESC
			LIMIT 1;
			`
	row := ssdb.db.QueryRow(sqlStmt)
	var endingState string
	err := row.Scan(&endingState)
	if err != nil {
		log.Fatal(err)
	}

	endingStateInt, err := strconv.Atoi(endingState)
	if err != nil {
		panic("unable to convert ending state")
	}
	return NewExhaustiveSearchState(int64(endingStateInt)) // We're assuming exhaustive search here, which we'll fix soon
}

func (ssdb *SqliteSearchDb) InsertSearchMetadata(smd SearchMetadata) {
	tx, err := ssdb.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt, err := tx.Prepare(`
		INSERT INTO
            SearchState(start_time, end_time, state_type, starting_state, ending_state)
            VALUES(?, ?, ?, ?, ?)
			;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()
	_, err = sqlStmt.Exec(
		smd.startTime.Format("2006-01-02 15:04:05-0700"),
		smd.endTime.Format("2006-01-02 15:04:05-0700"),
		smd.stateType,
		smd.startingState.Serialize(),
		smd.endingState.Serialize(),
	)

	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (ssdb *SqliteSearchDb) Close() {
	err := ssdb.db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Closed!")
}
