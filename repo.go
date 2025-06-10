package main

import (
	"database/sql"
	"fmt"
	"os"
)

type repo struct {
	db *sql.DB
}

func (s *server) newRepo() (*repo, error) {
	var r *repo
	var err error

	if err := createDBFile(DB_PATH); err != nil {
		return nil, err
	}

	r.db, err = sql.Open("sqlite3", DB_PATH)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	defer func() {
		if err != nil {
			r.db.Close()
		}
	}()

	if err := r.db.Ping(); err != nil {
		return nil, err
	}

	if err := r.createTables(); err != nil {
		return nil, err
	}

	return r, nil
}

func createDBFile(dbPath string) error {
	file, err := os.OpenFile(dbPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (r *repo) createTables() error {
	if _, err := r.db.Exec(INIT_MIGRATION); err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}
	return nil
}
