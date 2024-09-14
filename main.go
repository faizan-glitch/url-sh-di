package main

import (
	"log"
	"net/http"
	"url-sh/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.SetupRoutes(mux)

	http.ListenAndServe(":8080", mux)

	log.Println("pasidjaoisd")
}

// SERVER
// Shorten a long URL and Save it
// Check for collision
// Take Shortened URL and Return Original Long URL
