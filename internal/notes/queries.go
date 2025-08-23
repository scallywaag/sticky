package notes

const (
	CountNotesSQL = `SELECT COUNT(*) FROM notes;`

	ListNotesSQL = `
		SELECT id, content, color, status
		FROM notes;
	`

	AddNoteSQL = `
		INSERT INTO notes(id, content, color, status)
		VALUES(NULL, ?, ?, ?);
	`

	DeleteNoteSQL = `DELETE FROM notes WHERE id = ?;`

	GetMutationsSQL = `SELECT color, status FROM notes WHERE id = ?;`

	MutateNoteSQL = `
		UPDATE notes
		SET color = ?, status = ?
		WHERE id = ?;
	`
)
