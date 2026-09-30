package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Create
	// db.Create(&Product{
	// 	Name:  "Mouse",
	// 	Price: 40,
	// })

	// Buscando 1 PRODUTO * SELECT ONE *
	// var product Product
	// db.First(&product, 1)

	// Buscando com Where
	// db.First(&product, "name = ?", "Mouse")

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	// fmt.Println(product)

	// }

	// Delete
	// var productD Product
	// db.Where("id = ?", 6).Delete(&productD)
	// fmt.Println(productD)

	// // Deletando mais de um
	// db.Delete(&productD, []int{3, 4, 5, 7})

	// Aula de BelongsTo -> chaves estrangeiras
	// create category
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      40,
		CategoryID: category.ID,
	})

}
