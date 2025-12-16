package Handler

import (
	"encoding/json"
	"io"
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

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie := Models.Movie{
		Title:       r.FormValue("Title"),
		Description: r.FormValue("Description"),
		ReleaseYear: atoiSafe(r.FormValue("ReleaseYear")),
		Language:    r.FormValue("Language"),
		Duration:    atoiSafe(r.FormValue("Duration")),
		Rating:      r.FormValue("Rating"),
		AgeRating:   r.FormValue("AgeRating"),
		Country:     r.FormValue("Country"),
	}

	// Read files
	if file, _, _ := r.FormFile("Thumbnail"); file != nil {
		movie.Thumbnail, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("Banner"); file != nil {
		movie.Banner, _ = io.ReadAll(file)
		file.Close()
	}
	if file, _, _ := r.FormFile("Trailer"); file != nil {
		movie.Trailer, _ = io.ReadAll(file)
		file.Close()
	}
	if file, header, _ := r.FormFile("Movie"); file != nil {
		movie.MovieURL = header.Filename
		// optionally save movie file to disk
		file.Close()
	}

	if err := h.Service.CreateMovie(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Movie uploaded successfully",
	})
}

func atoiSafe(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (h *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(movies)
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
