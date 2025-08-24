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
		SELECT id FROM lists WHERE name = ?;
	`

	ListListsSQL = `
		WITH ordered_lists AS (
			SELECT
				ROW_NUMBER() OVER (ORDER BY id DESC) AS virtual_id,
				name
			FROM lists
		)
		SELECT virtual_id, name
		FROM ordered_lists;
	`

	AddListSQL = `
		INSERT INTO lists(name) VALUES(?);
	`

	DeleteListSQL = `
		WITH ordered_lists AS (
			SELECT
				ROW_NUMBER() OVER (ORDER BY id DESC) AS virtual_id,
				id
			FROM lists
		)
		DELETE FROM lists
		WHERE id = (SELECT id FROM ordered_lists WHERE virtual_id = ?);
	`
)
