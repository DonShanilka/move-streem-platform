package handlers

import (
    "encoding/json"
    // "fmt"
    "io"
    "net/http"
    "os"
    "github.com/DonShanilka/movie-service/internal/models"
    "github.com/DonShanilka/movie-service/internal/services"
    "strconv"
)

type MovieHandler struct {
    Service *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
    return &MovieHandler{Service: service}
}

// UploadMovie handles POST /upload
func (h *MovieHandler) UploadMovie(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    r.ParseMultipartForm(10 << 30) // 10 GB max file size

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "File error: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    title := r.FormValue("title")
    description := r.FormValue("description")
    genre := r.FormValue("genre")
    releaseYear, _ := strconv.Atoi(r.FormValue("release_year"))
    duration, _ := strconv.Atoi(r.FormValue("duration"))

    // Save file to local folder
    os.MkdirAll("./videos", os.ModePerm)
    filePath := "./videos/" + handler.Filename
    dst, _ := os.Create(filePath)
    defer dst.Close()
    io.Copy(dst, file)

    movie := models.Movie{
        Title:       title,
        Description: description,
        Genre:       genre,
        ReleaseYear: releaseYear,
        Duration:    duration,
        VideoURL:    "/videos/" + handler.Filename,
    }

    err = h.Service.SaveMovie(movie)
    if err != nil {
        http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message":   "Movie uploaded successfully",
        "file_name": handler.Filename,
    })
}

// ListMovies handles GET /movies
func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
    movies, err := h.Service.GetAllMovies()
    if err != nil {
        http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Return JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}
