package sqlite

import (
	"database/sql"
	notepkg "github.com/Ethanmurray3/notes-app/internal/note"
	"time"
)

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
func (r *Repository) Create(n notepkg.Note) (notepkg.Note, error) {
	now := time.Now()
	n.CreatedAt = now
	n.UpdatedAt = now

	query := `
	INSERT INTO notes (title, content, created_at, updated_at)
	VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, n.Title, n.Content, n.CreatedAt, n.UpdatedAt)
	if err != nil {
		return notepkg.Note{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return notepkg.Note{}, err
	}

	n.ID = id

	return n, nil
}
