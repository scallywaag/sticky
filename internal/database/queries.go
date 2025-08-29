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
			list_id INTEGER NOT NULL,
			FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
		);
	`

	StateSQL = `
		CREATE TABLE IF NOT EXISTS state (
			key TEXT PRIMARY KEY,
			list_id INTEGER,
			FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE SET NULL
		);
	`

	DefaultStateSQL = `
		INSERT INTO state(key, list_id)
		VALUES('active', NULL)
		ON CONFLICT (key) DO NOTHING;
	`
)
