package notes

import "database/sql"

type Repository interface {
	GetAll() ([]Note, error)
	Add(note *Note) error
	Delete(id int) error
	Update(id int) error
}

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) GetAll() ([]Note, error) {
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
