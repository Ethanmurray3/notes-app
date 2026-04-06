package main

import (
	"database/sql"
	"log"

	"github.com/Ethanmurray3/notes-app/internal/storage/sqlite"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "data/note.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	repository, err := sqlite.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	_ = repository
}
