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

		err = rows.Scan(&n.VirtualId, &n.Content, &n.Color, &n.Status)
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

func (r *DBRepository) Add(note *Note) error {
	return nil
}

func (r *DBRepository) Delete(id int) error {
	return nil
}

func (r *DBRepository) Update(id int) error {
	return nil
}

func (r *DBRepository) GetMutations(id int) error {
	return nil
}

func (r *DBRepository) Count(id int) error {
	return nil
}
