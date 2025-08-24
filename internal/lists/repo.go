package lists

import "database/sql"

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

	return nil, nil
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
	return 0, nil
}

func (r *DBRepository) GetId(name string) (int, error) {
	return 0, nil
}
