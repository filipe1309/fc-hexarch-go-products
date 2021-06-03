package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `
		CREATE TABLE products (
			id STRING,
			name STRING,
			price FLOAT,
			status STRING
		);
	`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "Prodcut Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}
