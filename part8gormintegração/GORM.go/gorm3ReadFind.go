// Modelos e Mapeamento
// GORM permite que você defina modelos usando estruturas Go.
// Você pode incluir campos especiais como gorm.Model para adicionar automaticamente
// campos como ID, CreatedAt, UpdatedAt e DeletedAt

package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Falid to connect database")
	}

	// Select Id key
	// var product Product
	//db.First(&product, 2) // id
	// db.First(&product, "name = ?", "Celular") // name
	// fmt.Println(product)

	// Selecionando todos os produtos
	//var products []Product
	//db.Find(&products)
	// db.Limit(2).Find(&products) // todos os produtos com limitaçao de "2"
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Consulta com WHERE, com mais detalhes
	var products []Product
	//db.Where("price > ?", 3500).Find(&products) // Por preço
	db.Where("name LIKE ?", "%Casa%").Find(&products) // Por nome ou letra
	for _, products := range products {
		fmt.Println(products)
	}
}
