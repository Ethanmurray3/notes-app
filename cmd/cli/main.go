package main

import (
	"database/sql"
	"github.com/Ethanmurray3/notes-app/internal/storage/sqlite"
	"log"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "note.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	defer db.Close()
	repository := sqlite.NewRepository(db)
	_ = repository
}
