package lists

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAll() ([]List, error)
	Add(name string) (int, error)
	Delete(id int) error
	GetActive() (*List, error)
	SetActive(id int, name string) (*List, error)
	Count() (int, error)
	GetId(name string) (int, error)
	GetFirst() (*List, error)
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

func (r *DBRepository) Add(name string) (int, error) {
	result, err := r.db.Exec(AddSQL, name)
	if err != nil {
		return 0, fmt.Errorf("exec failed: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return int(id), nil
}

func (r *DBRepository) Delete(virtualId int) error {
	result, err := r.db.Exec(DeleteSQL, virtualId)
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

func (r *DBRepository) GetActive() (*List, error) {
	l := &List{}
	err := r.db.QueryRow(GetActiveSQL).Scan(&l.Id, &l.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoActiveList
		}
		return nil, fmt.Errorf("query row failed: %w", err)
	}

	return l, nil
}

func (r *DBRepository) SetActive(id int, name string) (*List, error) {
	result, err := r.db.Exec(SetActiveSQL, id, id)
	if err != nil {
		return nil, fmt.Errorf("failed to set active list: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return nil, ErrNoRowsAffected
	}

	l := &List{
		Id:   id,
		Name: name,
	}
	return l, nil
}

func (r *DBRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow(CountSQL).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("query row failed: %w", err)
	}

	return count, nil
}

func (r *DBRepository) GetId(name string) (int, error) {
	var id int
	err := r.db.QueryRow(GetIdByNameSQL, name).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("list %q does not exist", name)
		}
		return 0, fmt.Errorf("failed to query list: %w", err)
	}

	return id, nil
}

func (r *DBRepository) GetFirst() (*List, error) {
	var l List
	err := r.db.QueryRow("SELECT id, name FROM lists ORDER BY id LIMIT 1").Scan(&l.Id, &l.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no lists found")
		}
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return &l, nil
}
