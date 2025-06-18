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

	//Atualizando um produto com base no seu id
	var p Product
	db.First(&p, 1)
	p.Name = "Data Show"
	db.Save(&p)

	var products []Product
	//db.Where("price > ?", 3500).Find(&products)       // Por preço
	db.Where("name LIKE ?", "%Data%").Find(&products) // Por nome ou letra
	for _, products := range products {
		fmt.Println(products)
	}

	//Atualizando Múltiplas Colunas
	db.Model(&p).Updates(Product{Price: 6000, Name: "Xbox Series X"})
	//db.Model(&Product{}).Where("id = ?", 1).Updates(Product{Price: 9999, Name: "Bicicleta"})

	//Exclusão
	var p3 Product
	db.First(&p3, 4) // por id
	db.Delete(&p3)

}
