package main

import (
	"appproduct/adapters/db"
	"appproduct/application"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db2, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db.NewProductDb(db2)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}
