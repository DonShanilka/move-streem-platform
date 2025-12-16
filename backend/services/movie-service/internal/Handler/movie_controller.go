package Handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DonShanilka/movie-service/internal/Models"
	services "github.com/DonShanilka/movie-service/internal/Service"
)

type MovieHandler struct {
	Service *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
	return &MovieHandler{Service: service}
}

func (h *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	if err := h.Service.CreateMovie(&movie); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Movie created"})
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var movie Models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	if err := h.Service.UpdateMovie(uint(id), &movie); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Movie updated"})
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if err := h.Service.DeleteMovie(uint(id)); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted"})
}
