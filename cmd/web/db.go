package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDb(dbPath string) (*sql.DB, error) {
	dbConnection := fmt.Sprintf("%s?_loc=auto", dbPath)
	db, err := sql.Open("sqlite3", dbConnection)

	if err != nil {
		return nil, err
	}

	return db, nil
}
