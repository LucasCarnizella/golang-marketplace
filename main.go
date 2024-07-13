package main

import (
	"github.com/joho/godotenv"
	"log"
	"marketplace/routes"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	appUrl := os.Getenv("APP_URL")

	// Call function to open the APP URL in the browser
	err = openUrl(appUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Load routes and start server
	routes.LoadRoutes()
	err = http.ListenAndServe(appUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func openUrl(targetUrl string) error {
	var cmd string
	var args []string

	// Check the existence of the URL scheme, if not assign the default scheme "http"
	if !strings.HasSuffix(targetUrl, "https://") && !strings.HasSuffix(targetUrl, "http://") {
		_, err := url.Parse("http://" + targetUrl)
		if err != nil {
			log.Fatal(err)
		}

		targetUrl = "http://" + targetUrl
	}

	// Decide which command to use depending on the Host OS
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}

	args = append(args, targetUrl)

	return exec.Command(cmd, args...).Start()
}
