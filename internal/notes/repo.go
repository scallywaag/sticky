package notes

import "database/sql"

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

func (r *DBRepository) GetAll(name string) ([]Note, error) {
	return nil, nil
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
