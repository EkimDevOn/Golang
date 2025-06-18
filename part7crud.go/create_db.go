// ATUALIZAÇÃO que nosso ID tenha um id mais robusto
// precisamos do pacote install terminal root project:
// go get github.com/google/uuid - Uma ferramenta útil para gerar e manipular UUIDs de maneira eficiente e segura em seus projetos Go.
// criamaos um novo banco de dados e uma nova tabela
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    string
	Name  string
	Email string
	Age   int
}

func NewUser(name, email string, age int) *User {
	return &User{
		Id:    uuid.New().String(),
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func creatTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER
	);
	`

	_, err := db.Exec(query)
	return err
}

func insertUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("INSERT INTO users(id, name, email, age) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	return nil
}

// update user
func updateUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.Id)
	if err != nil {
		return err
	}
	return nil
}

// Select user
func selectUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u) // mapeando users
	}
	return users, nil
}

// Delete users
func deleteUser(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id) // Passar o ID como argumento
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela
	err = creatTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// Inserção de Usuário
	user := NewUser("Jet Li", "KungFu@example.com", 25)
	err = insertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	//Atualização de usuário
	user.Name = "Jhon Smith"
	user.Age = 35
	err = updateUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	// Listagem de usuários
	users, err := selectUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("Usuário %s possui email %s\n", u.Name, u.Email)
	}

	// Exclusão de dados
	err = deleteUser(db, user.Id)
	 if err != nil {
	 	log.Fatal(err)
	}
}