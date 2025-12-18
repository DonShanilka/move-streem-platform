package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/genres-service/internal/Handler"
	"github.com/DonShanilka/genres-service/internal/Repository"
	"github.com/DonShanilka/genres-service/internal/Routes"
	"github.com/DonShanilka/genres-service/internal/Service"
	"github.com/DonShanilka/genres-service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	genreRepo := Repository.NewGenerRepostry(database)
	genreService := Service.NewGenreService(genreRepo)
	genreHandler := Handler.NewGenreHandler(genreService)

	mux := http.NewServeMux()
	Routes.RegisterGenreRoutes(mux, genreHandler)

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if request.Method == http.MethodOptions {
			return
		}
		mux.ServeHTTP(writer, request)
	})

	log.Println("Genres Service running on :8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
