package Handler

import (
	"encoding/json"
	//"io"
	"net/http"
	//"strconv"

	"github.com/DonShanilka/genres-service/internal/Models"
	"github.com/DonShanilka/genres-service/internal/Service"
)

type GenreHandler struct {
	Service *Service.GenreService
}

func NewGenreHandler(service *Service.GenreService) *GenreHandler {
	return &GenreHandler{Service: service}
}

func (handler *GenreHandler) CreateGenre(writer http.ResponseWriter, request *http.Request) {

	var genre Models.Genre

	if err := json.NewDecoder(request.Body).Decode(&genre); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if genre.Name == "" {
		http.Error(writer, "Genre name is required", http.StatusBadRequest)
		return
	}

	err := handler.Service.CreateGenre(&genre)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(genre)

}
