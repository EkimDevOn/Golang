package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db") // sql.Opem() 2x aplicaçoes cria banco de dados se nao existe e se existe conecta no DB
	if err != nil {
		log.Fatal("Error connection DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users(name) VALUES(?)", "Bruna Marquês", 2) // Dados diretamente na cláusula não é  a melhor prática por questoês de segurança, (SQL INJECTION)
	if err != nil {
		log.Fatal("Error to add values in table", err)
	}
	log.Println("Novo usuário criado com sucesso!")
}
