package sqlite

import "database/sql"

type Repository struct {
	db *sql.DB
}

func (r *Repository) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	`

	_, err := r.db.Exec(query)
	return err
}
