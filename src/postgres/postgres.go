// Package postgres is a simple implementation of postgres storage.
package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Postgres contain database data to run commands.
type Postgres struct {
	db *sql.DB
}

// InitializeDatabase gets the database connection information and returns a new Postgres instance.
func InitializeDatabase(connectionSting string) (*Postgres, error) {
	pgDB, err := sql.Open("postgres", connectionSting)
	if err != nil {
		return nil, err
	}
	postgres := &Postgres{}
	postgres.db = pgDB
	return postgres, nil
}

// ExecQuery gets a postgres query and returns result as sql rows.
func (postgres *Postgres) ExecQuery(query string) (*sql.Rows, error) {
	rows, err := postgres.db.Query(query)
	return rows, err

}

// ExecQuery gets a postgres query and returns result as sql rows.
func (postgres *Postgres) InsertRandomDataIntoTrade() error {
	_, err := postgres.db.Query(`INSERT INTO TRADE(ID,

		INSTRUMENTID,
		DATEEN,
		OPEN,
		HIGH,
		LOW,
		CLOSE)
SELECT FLOOR(RANDOM() * 1000 + 1)::int,
FLOOR(RANDOM() * 1000 + 1)::int,
TIMESTAMP '2020-01-10 20:00:00' + RANDOM() * (TIMESTAMP '2021-01-20 00:00:00' - TIMESTAMP '2020-01-10 10:00:00'),
FLOOR(RANDOM() * 10000 + 1)::int,
FLOOR(RANDOM() * 10000 + 1)::int,
FLOOR(RANDOM() * 10000 + 1)::int,
FLOOR(RANDOM() * 10000 + 1)::int;`)
	return err

}
