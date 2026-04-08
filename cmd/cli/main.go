package main

import (
	"database/sql"
	notepkg "github.com/Ethanmurray3/notes-app/internal/note"
	sqlitepkg "github.com/Ethanmurray3/notes-app/internal/storage/sqlite"
	"log"
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

	repository, err := sqlitepkg.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	_ = repository

	note := notepkg.Note{
		Title:   "First Note",
		Content: "Hello from Go",
	}

	created, err := repository.Create(note)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created note:", created)
}
