package main

const (
	DB_PATH        = `pekarica.db`
	INIT_MIGRATION = `
	CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		instructions TEXT,
		prep_time INTEGER,
		cook_time INTEGER,
		servings INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS ingredients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		unit TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS recipe_ingredients (
		recipe_id INTEGER NOT NULL,
		ingredient_id INTEGER NOT NULL,
		quantity REAL,
		notes TEXT,
		PRIMARY KEY (recipe_id, ingredient_id),
		FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
		FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		recipe_id INTEGER NOT NULL,
		date_made TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		notes TEXT,
		rating INTEGER CHECK (rating BETWEEN 1 AND 5),
		FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
	);
	`
)
