package models

import (
	"database/sql"
	"log"
	"marketplace/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SelectAllProducts() []Product {
	dbConn := db.Connection()

	productsQuery, err := dbConn.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	p := Product{}
	var products []Product

	for productsQuery.Next() {
		var id int
		var name string
		var description string
		var price float64
		var amount int

		err = productsQuery.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			log.Fatal(err)
		}

		p.ID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)

	return products
}
