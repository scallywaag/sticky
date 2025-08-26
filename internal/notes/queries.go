package notes

const (
	CountSQL = `SELECT COUNT(*) FROM notes WHERE list_id = ?;`

	GetAllSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin' THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id DESC
				) AS virtual_id,
				content,
				color,
				status
			FROM notes
			WHERE list_id = ?
		)
		SELECT virtual_id, content, color, status
		FROM ordered_notes;
	`

	AddSQL = `
		INSERT INTO notes(content, color, status, list_id)
		VALUES(?, ?, ?, ?);
	`

	DeleteSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin' THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id DESC
				) AS virtual_id,
				id
			FROM notes
			WHERE list_id = ?
		)
		DELETE FROM notes
		WHERE id = (SELECT id FROM ordered_notes WHERE virtual_id = ?);
	`

	GetMutationsSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin'   THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id DESC
				) AS virtual_id,
				color,
				status
			FROM notes
			WHERE list_id = ?
		)
		SELECT color, status
		FROM ordered_notes
		WHERE virtual_id = ?;
	`

	UpdateSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin' THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id DESC
				) AS virtual_id,
				id,
				color,
				status
			FROM notes
			WHERE list_id = ?
		)
		UPDATE notes
		SET color = ?, status = ?
		WHERE id = (SELECT id FROM ordered_notes WHERE virtual_id = ?);
	`
)
