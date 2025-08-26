package notes

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAll(name string) ([]Note, error)
	Add(note *Note) error
	Delete(id int) error
	Update(id int) error
	GetMutations(id int) error
	Count(id int) (int, error)
}

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) GetAll(activeListId string) ([]Note, error) {
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
		return fmt.Errorf("delete failed: no rows affected")
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

func (r *DBRepository) GetMutations(id int) error {
	return nil
}

func (r *DBRepository) Count(id int) (int, error) {
	var count int
	err := r.db.QueryRow(CountSQL, id).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("query row failed: %w", err)
	}
	return count, nil
}
