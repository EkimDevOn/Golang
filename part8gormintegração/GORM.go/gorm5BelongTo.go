// determinado item pertence a outro
// contexto produtos relacionamentos de produto com categoria
// estrtuturando este tipo de relacionamento

package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	Id   int `gorm: primaryKey`
	Name string
}

type Product struct {
	Id         int `gorm: primaryKey`
	Name       string
	Code       string
	Price      float64
	CategoryId int      // Chave estrangeira que referencia o Category
	Category   Category // Relacionamento "Belongs To"
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Falid to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Inserindo Categoria)
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// inserindo produto
	db.Create(&Product{
		Name:       "Fogão",
		Price:      1500.00,
		CategoryId: category.Id,
	})

	// listando Produtos
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, "|", product.Category.Name, "|", product.Price)
	}
}
