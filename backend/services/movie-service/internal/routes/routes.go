package routes

import (
    "net/http"
    "github.com/DonShanilka/movie-service/internal/handlers"
)

func MovieRoutes(mux *http.ServeMux, h *handlers.MovieHandler) {
    mux.HandleFunc("/api/movies", h.ListMovies)
    mux.HandleFunc("/api/movies/upload", h.UploadMovie)
    mux.HandleFunc("/api/movies/stream", h.StreamMovie)
}


func SeriesRoutes(mux *http.ServeMux, h *handlers.SeriesHandler) {
    // mux.HandleFunc("/api/series", h.ListSeries)
    mux.HandleFunc("/api/series/upload", h.UpdaloadSeries)
    // mux.HandleFunc("/api/series/stream", h.StreamSeries)
}

func EpisodeRouts(mux *http.ServeMux, h *handlers.EpisodeHandler) {
    mux.HandleFunc("/api/episodes/upload", h.UploadEpisode)
}
