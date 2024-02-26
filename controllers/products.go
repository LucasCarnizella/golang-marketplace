package controllers

import (
	"html/template"
	"log"
	"marketplace/models"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectAllProducts()

	err := templates.ExecuteTemplate(w, "Index", products)
	if err != nil {
		log.Fatal(err)
	}
}
