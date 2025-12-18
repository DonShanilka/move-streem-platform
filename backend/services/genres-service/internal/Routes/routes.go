package Routes

import (
	"net/http"

	"github.com/DonShanilka/genres-service/internal/Handler"
)

func RegisterGenreRoutes(mux *http.ServeMux, handler *Handler.GenreHandler) {
	mux.HandleFunc("/api/genre/creatGenres", handler.CreateGenre)
	//mux.HandleFunc("/api/movies/createMovie", h.CreateMovie)
	//mux.HandleFunc("/api/movies/updateMovie", h.UpdateMovie)
	//mux.HandleFunc("/api/movies/deleteMovie", h.DeleteMovie)
}
