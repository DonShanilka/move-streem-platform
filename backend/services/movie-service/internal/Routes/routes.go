package Routes

import (
	"net/http"

	"github.com/DonShanilka/movie-service/internal/Handler"
)

func RegisterMovieRoutes(mux *http.ServeMux, h *Handler.MovieHandler) {
	mux.HandleFunc("/api/movies/getAllMovies", h.GetAllMovies)
	mux.HandleFunc("/api/movies/create", h.CreateMovie)
	mux.HandleFunc("/api/movies/update", h.UpdateMovie)
	mux.HandleFunc("/api/movies/delete", h.DeleteMovie)
}
