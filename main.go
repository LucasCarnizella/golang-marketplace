package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"text/template"

	_ "github.com/lib/pq"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	appHost := os.Getenv("APP_HOST")

	db := dbConnection()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// To-Do handle http & https urls
	err = openUrl("http://" + appHost)

	http.HandleFunc("/", index)
	err = http.ListenAndServe(appHost, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()

	selectAllProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	p := Product{}
	var products []Product

	for selectAllProducts.Next() {
		var id int
		var name string
		var description string
		var price float64
		var amount int

		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)
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

	err = templates.ExecuteTemplate(w, "Index", products)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}

func dbConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := "dbname=" + dbName + " user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func openUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
