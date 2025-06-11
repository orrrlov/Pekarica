package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db *sql.DB
}

func (s *server) newRepo() (*repo, error) {
	var err error
	r := &repo{}

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

	r.db.SetMaxOpenConns(1)
	r.db.SetMaxIdleConns(1)
	r.db.SetConnMaxLifetime(0)

	if err := r.createTables(); err != nil {
		return nil, err
	}

	return r, nil
}

func createDBFile(dbPath string) error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		return file.Close()
	}
	return nil
}

func (r *repo) createTables() error {
	if _, err := r.db.Exec(INIT_MIGRATION); err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}
	return nil
}
