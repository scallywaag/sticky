package lists

const (
	CountListsSQL = `SELECT COUNT(*) FROM lists;`

	GetActiveListSQL = `
		SELECT l.id, l.name
	    FROM state s
	    JOIN lists l ON s.list_id = l.id
	    WHERE s.key = 'active';
	`

	SetActiveListSQL = `
        UPDATE state
        SET list_id = ?
        WHERE key = 'active';
	`

	GetListIdByNameSQL = `
		SELECT id FROM lists WHERE name = ?
	`

	ListListsSQL = `
		SELECT id, name FROM lists;
	`
	AddListSQL = `
	`

	DeleteListSQL = ``
)
