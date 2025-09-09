package notes

import (
	"database/sql"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

type Repository interface {
	GetAll(activeListId int) ([]Note, error)
	Add(note *Note, activeListId int) error
	Delete(id, activeListId int) error
	Update(note *Note, activeListId int) error
	GetMutations(id, activeListId int) (formatter.Color, NoteStatus, error)
	Count(id int) (int, error)
	CheckNotesExist(activeListId int) (bool, error)
	CheckNoteExists(id, activeListId int) (bool, error)
}

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) GetAll(activeListId int) ([]Note, error) {
	rows, err := r.db.Query(GetAllSQL, activeListId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		n := Note{}

		err = rows.Scan(&n.Id, &n.Content, &n.Color, &n.Status)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return notes, nil
}

func (r *DBRepository) Add(note *Note, activeListId int) error {
	_, err := r.db.Exec(AddSQL, note.Content, note.Color, note.Status, activeListId)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	return nil
}

func (r *DBRepository) Delete(id int, activeListId int) error {
	result, err := r.db.Exec(DeleteSQL, activeListId, id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNoRowsAffected
	}

	return nil
}

func (r *DBRepository) Update(note *Note, activeListId int) error {
	result, err := r.db.Exec(UpdateSQL, activeListId, note.Color, note.Status, note.Id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("delete failed: no rows affected")
	}

	return nil
}

func (r *DBRepository) GetMutations(id, activeListId int) (formatter.Color, NoteStatus, error) {
	var currentColor formatter.Color
	var currentStatus NoteStatus

	err := r.db.QueryRow(GetMutationsSQL, activeListId, id).
		Scan(&currentColor, &currentStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", ErrNoRowsResult
		}
		return "", "", fmt.Errorf("query row failed: %w", err)
	}

	return currentColor, currentStatus, nil
}

func (r *DBRepository) Count(id int) (int, error) {
	var count int
	err := r.db.QueryRow(CountSQL, id).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("query row failed: %w", err)
	}
	return count, nil
}

func (r *DBRepository) CheckNotesExist(activeListId int) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM notes
			WHERE list_id = ?
		)
	`
	err := r.db.QueryRow(query, activeListId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("query row failed: %w", err)
	}
	return exists, nil
}

func (r *DBRepository) CheckNoteExists(id, activeListId int) (bool, error) {
	var exists bool
	query := `
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
				) AS virtual_id
			FROM notes
			WHERE list_id = ?
		)
		SELECT EXISTS (
			SELECT 1
			FROM ordered_notes
			WHERE virtual_id = ?
		);
	`
	err := r.db.QueryRow(query, activeListId, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("query row failed: %w", err)
	}
	return exists, nil
}
