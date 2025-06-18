package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	Id       int `gorm: primaryKey`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	Id         int `gorm: primaryKey`
	Name       string
	Code       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Falid to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})

	// // Inserindo Categoria
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// category2 := Category{Name: "Eletrodomésticos"}
	// db.Create(&category2)

	// // inserindo produto
	// db.Create(&Product{
	// 	Name:  "mesa",
	// 	Price: 2500.00,
	// 	//CategoryId: 1,
	// 	Categories: []Category{category, category2},
	// })

	//Listando categorias e Produtos
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("->", product.Name)
		}
	}
}
