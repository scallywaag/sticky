package lists

const (
	CountSQL = `SELECT COUNT(*) FROM lists;`

	GetActiveSQL = `
		SELECT l.id, l.name
		FROM state s
		JOIN lists l ON s.list_id = l.id
		WHERE s.key = 'active';
	`

	SetActiveSQL = `
        UPDATE state
        SET list_id = ?
        WHERE key = 'active'
		AND EXISTS (SELECT 1 FROM lists WHERE id = ?);
	`

	GetIdByNameSQL = `
		SELECT id FROM lists WHERE name = ?;
	`

	GetAllSQL = `
		WITH ordered_lists AS (
			SELECT
				ROW_NUMBER() OVER (ORDER BY id DESC) AS virtual_id,
				name
			FROM lists
		)
		SELECT virtual_id, name
		FROM ordered_lists;
	`

	AddSQL = `
		INSERT INTO lists(name) VALUES(?);
	`

	DeleteSQL = `
		WITH ordered_lists AS (
			SELECT
				ROW_NUMBER() OVER (ORDER BY id DESC) AS virtual_id,
				id
			FROM lists
		)
		DELETE FROM lists
		WHERE id = (SELECT id FROM ordered_lists WHERE virtual_id = ?);
	`

	GetFirstListSQL = `
		SELECT id, name FROM lists ORDER BY id LIMIT 1;
	`
)
