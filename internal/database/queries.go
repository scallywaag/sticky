package database

const (
	ListsSQL = `
		CREATE TABLE IF NOT EXISTS lists (
			id INTEGER PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		);
	`

	NotesSQL = `
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY,
			content TEXT NOT NULL,
			color TEXT,
			status TEXT NOT NULL DEFAULT 'default',
			list_id INTEGER NOT NULL DEFAULT 1,
			FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
		);
	`

	StateSQL = `
		CREATE TABLE IF NOT EXISTS state (
			key TEXT PRIMARY KEY,
			list_id INTEGER NOT NULL,
			FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
		);
	`
)
