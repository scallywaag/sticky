package notes

const (
	CountNotesSQL = `SELECT COUNT(*) FROM notes;`

	ListNotesSQL = `
		SELECT id, content, color, cross
		FROM notes;
	`

	AddNoteSQL = `
		INSERT INTO notes(id, content, color, cross)
		VALUES(NULL, ?, ?, ?);
	`

	DeleteNoteSQL = `DELETE FROM notes WHERE id = ?;`

	GetMutationsSQL = `SELECT color, cross FROM notes WHERE id = ?;`

	MutateNoteSQL = `
		UPDATE notes
		SET color = ?, cross = ?
		WHERE id = ?;
	`
)
