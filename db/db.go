package db

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDB(url string) *sql.DB {
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	}

	return db
}
