package main

import (
	"database/sql"
	"fmt"
	"math"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	pageSize = 20
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

func (r *repo) getAllIngredients() ([]Ingredient, error) {
	var is []Ingredient
	query := `
	SELECT id, name, description 
	FROM ingredients 
	ORDER BY id DESC;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i Ingredient
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

func (r *repo) createIngredient(i Ingredient) error {
	query := `
	INSERT INTO ingredients (name, description) 
	VALUES (?, ?);`
	if _, err := r.db.Exec(query, i.Name, i.Description); err != nil {
		return err
	}
	return nil
}

func (r *repo) deleteIngredient(id string) error {
	query := `
	DELETE FROM ingredients 
	WHERE id = ?;`
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (r *repo) getRecipesPagination() int {
	var totalCount int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM recipes`).Scan(&totalCount)
	if err != nil {
		return 0
	}
	return int(math.Ceil(float64(totalCount) / float64(pageSize)))
}

func (r *repo) getAllRecipes(page int) ([]Recipe, error) {
	var rps []Recipe

	offset := (page - 1) * pageSize

	query := `
	SELECT id,name,description,quantity,unit
    FROM recipes
    ORDER BY id DESC
    LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rp Recipe

		if err := rows.Scan(
			&rp.ID,
			&rp.Name,
			&rp.Description,
			&rp.Quantity,
			&rp.Unit,
		); err != nil {
			return nil, err
		}

		rps = append(rps, rp)
	}

	return rps, nil
}

func (r *repo) createRecipe(rp Recipe) error {
	query := `
	INSERT INTO recipes (name, description, quantity, unit) 
	VALUES (?, ?, ?, ?);`

	if _, err := r.db.Exec(query, rp.Name, rp.Description, rp.Quantity, rp.Unit); err != nil {
		return err
	}

	return nil
}
