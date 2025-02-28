package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const ()

func New(
	host string,
	port int,
	user string,
	password string,
	dbname string,

) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
