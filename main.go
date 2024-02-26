package main

import (
	"github.com/joho/godotenv"
	"log"
	"marketplace/routes"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	appHost := os.Getenv("APP_HOST")

	// To-Do handle http & https urls
	err = openUrl("http://" + appHost)

	routes.LoadRoutes()
	err = http.ListenAndServe(appHost, nil)
	if err != nil {
		log.Fatal(err)
	}
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
