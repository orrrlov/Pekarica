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

func (r *repo) getAllIngredients() ([]ingredient, error) {
	var is []ingredient
	query := `
	SELECT id, name, description 
	FROM ingredients 
	ORDER id DESC;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i ingredient
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		is = append(is, i)
	}
	return is, nil
}

func (r *repo) createIngredient(i ingredient) error {
	query := `
	INSERT INTO ingredients (name, description) 
	VALUES (?, ?);`
	if _, err := r.db.Exec(query, i.Name, i.Description); err != nil {
		return err
	}
	return nil
}
