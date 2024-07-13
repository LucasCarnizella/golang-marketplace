package controllers

import (
	"html/template"
	"log"
	"marketplace/models"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	err := templates.ExecuteTemplate(w, "Index", products)
	if err != nil {
		log.Fatal(err)
	}
}

func NewProductPage(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "NewProduct", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func EditProductPage(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	product := models.GetProduct(ID)

	err := templates.ExecuteTemplate(w, "EditProduct", product)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal(err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal(err)
		}

		models.CreateProduct(name, description, convertedPrice, convertedAmount)
	}

	http.Redirect(w, r, "/", 301)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal(err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal(err)
		}

		models.UpdateProduct(convertedID, name, description, convertedPrice, convertedAmount)
	}

	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	models.DeleteProduct(ID)

	http.Redirect(w, r, "/", 301)
}
