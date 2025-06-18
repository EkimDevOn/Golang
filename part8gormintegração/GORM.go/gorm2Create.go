// GORM é um framework ORM (Object-Relational Mapping) para a linguagem Go.
// Instalação
//go mod init go-gorm  // cria modulo novo
//go mod tidy
// Para começar a usar GORM, você precisa instalar o pacote e o driver do banco de dados que você pretende usar:
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/sqlite
// go get -u gorm.io/driver/mysql
// go get -u gorm.io/driver/postgres
// go get -u gorm.io/driver/sqlserver

package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model // soft delete
	Name       string
	Code       string
	Price      float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Falid to connect database")
	}

	// Migrar schema
	db.AutoMigrate(&Product{})

	// Criar
	//db.Create(&Product{Code: "D44", Price: 25})
	products := []Product{
		{Name: "Nootbook", Code: "D45", Price: 5000.00},
		{Name: "Celular", Code: "D46", Price: 4500.00},
		{Name: "Tablet", Code: "D47", Price: 2300.00},
		{Name: "Carro", Code: "D48", Price: 550000.00},
		{Name: "Moto", Code: "D49", Price: 50000.00},
		{Name: "Casa", Code: "D50", Price: 1500000.00},
	}
	db.Create(&products)
}
