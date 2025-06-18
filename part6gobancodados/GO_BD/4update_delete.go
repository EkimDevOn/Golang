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
	// Atualizando dados func.Exec()
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Mateo", 10) // passando a funçao diretamente, caso contrário complexos, use variáveis
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Updated %d row(s)\n", affectedRows)

	//Eclusão de Dados
	result, err = db.Exec("DELETE FROM users WHERE id = ?", 10) // no caso o id
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
}

// ps: Nao esqueça de atualizar o mod dity manualmente na sua "file go mod",
// ou comando no terminal-root-project: run go mod tidy
