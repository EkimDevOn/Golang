package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Error connection DB", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("Err")
	}
	defer rows.Close()

	for rows.Next() { // na variável via *ponteiro
		var id int
		var name string
		err = rows.Scan(&id, &name) // primeiro id-int, name-string
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("User: %d -> %s", id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
