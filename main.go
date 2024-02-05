package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := dbConnection()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	//	products := []Product{
	//		{Name: "T-Shirt", Description: "Black", Price: 49.99, Quantity: 30},
	//		{Name: "Snicker", Description: "Size 20", Price: 109.99, Quantity: 10},
	//	}

	templates.ExecuteTemplate(w, "Index", nil)
}

func dbConnection() *sql.DB {
	connStr := "dbname=marketplace user=postgres password=inova2023 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
