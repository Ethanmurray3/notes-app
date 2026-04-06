package sqlite

import "database/sql"

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
