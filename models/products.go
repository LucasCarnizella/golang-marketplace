package models

import (
	"database/sql"
	"log"
	"marketplace/db"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Amount      int
	CreatedAt   time.Time
}

func GetProduct(ID string) Product {
	dbConn := db.Connection()

	productQuery, err := dbConn.Query("SELECT * FROM products WHERE id=$1", ID)
	if err != nil {
		log.Fatal(err)
	}

	products := synthesizeProducts(productQuery)
	product := products[0]

	return product
}

func GetAllProducts() []Product {
	dbConn := db.Connection()

	productsQuery, err := dbConn.Query("SELECT * FROM products ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	products := synthesizeProducts(productsQuery)

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)

	return products
}

func CreateProduct(Name string, Description string, Price float64, Amount int) {
	dbConn := db.Connection()

	insertData, err := dbConn.Prepare("INSERT INTO products (name, description, price, amount) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = insertData.Exec(Name, Description, Price, Amount)
	if err != nil {
		log.Fatal(err)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)
}

func UpdateProduct(ID int, Name string, Description string, Price float64, Amount int) {
	dbConn := db.Connection()

	updateData, err := dbConn.Prepare("UPDATE products SET name=$2, description=$3, price=$4, amount=$5 WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = updateData.Exec(ID, Name, Description, Price, Amount)
	if err != nil {
		log.Fatal(err)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)
}

func DeleteProduct(ID string) {
	dbConn := db.Connection()

	deleteData, err := dbConn.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = deleteData.Exec(ID)
	if err != nil {
		log.Fatal(err)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)
}

func synthesizeProducts(query *sql.Rows) []Product {
	product := Product{}
	var products []Product

	for query.Next() {
		var id int
		var name string
		var description string
		var price float64
		var amount int
		var createdAt time.Time

		err := query.Scan(&id, &name, &description, &price, &amount, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
		product.CreatedAt = createdAt

		products = append(products, product)
	}

	return products
}
