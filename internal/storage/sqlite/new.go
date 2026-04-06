package sqlite

import "database/sql"

func NewRepository(db *sql.DB) (*Repository, error) {
	repo := &Repository{db: db}

	err := repo.migrate()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
