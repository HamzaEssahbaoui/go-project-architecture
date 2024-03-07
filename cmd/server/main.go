package main

import (
	"best-architecture/internal/config" // Add missing import
	"best-architecture/internal/handlers"
	"best-architecture/internal/services"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load(".env")

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	movieService := services.NewMovieDBService(cfg.APIKey, &http.Client{})
	movieHandler := handlers.NewMovieHandler(movieService)
	homeHandler := handlers.NewHomeHandler(movieService)

	http.HandleFunc("/", homeHandler.ServeHTTP)
	http.HandleFunc("/movie/", movieHandler.ServeHTTP)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
