package main

import (
    "log"
    "net/http"
    "github.com/DonShanilka/movie-service/internal/handlers"
    "github.com/DonShanilka/movie-service/internal/repository"
    "github.com/DonShanilka/movie-service/internal/routes"
    "github.com/DonShanilka/movie-service/internal/services"
)

func main() {
    // Initialize DB (MySQL)
    db, err := repository.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    // Initialize services and handlers
    movieService := services.NewMovieService(db)
    movieHandler := handlers.NewMovieHandler(movieService)

    // Setup routes
    handler := routes.SetupRoutes(movieHandler)

    log.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
