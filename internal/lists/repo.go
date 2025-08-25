package lists

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAll() ([]List, error)
	Add(name string) error
	Delete(id int) error
	GetActive() (*List, error)
	SetActive(name string) (*List, error)
	Count() (int, error)
	GetId(name string) (int, error)
}

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) GetAll() ([]List, error) {
	rows, err := r.db.Query(GetAllSQL)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	lists := []List{}
	for rows.Next() {
		var l List
		if err := rows.Scan(&l.Id, &l.Name); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		lists = append(lists, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return lists, nil
}

func (r *DBRepository) Add(name string) error {
	return nil
}

func (r *DBRepository) Delete(id int) error {
	return nil
}

func (r *DBRepository) GetActive() (*List, error) {
	return nil, nil
}

func (r *DBRepository) SetActive(name string) (*List, error) {
	return nil, nil
}

func (r *DBRepository) Count(name string) (int, error) {
	var count int
	err := r.db.QueryRow(CountSQL).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("query	 row failed: %w", err)
	}

	return count, nil
}

func (r *DBRepository) GetId(name string) (int, error) {
	return 0, nil
}
