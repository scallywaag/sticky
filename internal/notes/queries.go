package notes

const (
	CountNotesSQL = `SELECT COUNT(*) FROM notes;`

	ListNotesSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin' THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id
				) AS virtual_id,
				content,
				color,
				status
			FROM notes
		)
		SELECT virtual_id, content, color, status
		FROM ordered_notes
	`
	AddNoteSQL = `
		INSERT INTO notes(id, content, color, status)
		VALUES(NULL, ?, ?, ?);
	`

	DeleteNoteSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin' THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id
				) AS virtual_id,
				id
			FROM notes
		)
		DELETE FROM notes
		WHERE id = (SELECT id FROM ordered_notes WHERE virtual_id = ?)
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
				id
			) AS virtual_id,
			color,
			status
		  FROM notes
		)
		SELECT color, status
		FROM ordered_notes
		WHERE virtual_id = ?;
	`

	MutateNoteSQL = `
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN status = 'pin' THEN 1
							WHEN status = 'cross' THEN 3
							ELSE 2
						END,
						id
				) AS virtual_id,
				id,
				color,
				status
			FROM notes
		)
		UPDATE notes
		SET color = ?, status = ?
		WHERE id = (SELECT id FROM ordered_notes WHERE virtual_id = ?)
	`
)
