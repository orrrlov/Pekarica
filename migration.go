package main

const (
	DB_PATH        = `pekarica.db`
	INIT_MIGRATION = `
	-- Enable foreign key constraints
	PRAGMA foreign_keys = ON;

	-- Recipes table
	CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		quantity REAL,
		unit TEXT
	);

	-- Ingredients table
	CREATE TABLE IF NOT EXISTS ingredients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT
	);

	-- Recipe-Ingredients junction table (with amounts)
	CREATE TABLE IF NOT EXISTS recipe_ingredients (
		recipe_id INTEGER NOT NULL,
		ingredient_id INTEGER NOT NULL,
		quantity REAL NOT NULL,
		unit TEXT NOT NULL,
		PRIMARY KEY (recipe_id, ingredient_id),
		FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
		FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE
	);
	`
)
