package Handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
)

type TvSeriesHandler struct {
	Service *Service.TvSerriesService
}

func NewTvSeriesHandler(service *Service.TvSerriesService) *TvSeriesHandler {
	return &TvSeriesHandler{Service: service}
}

func (h *TvSeriesHandler) CreateTvSeries(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tvSeries := Models.Series{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ReleaseYear: atoiSafe(r.FormValue("releaseYear")),
		SeasonCount: atoiSafe(r.FormValue("seasonCount")),
		Language:    r.FormValue("language"),
	}

	if file, _, _ := r.FormFile("banner"); file != nil {
		tvSeries.Banner, _ = io.ReadAll(file)
		file.Close()
	}

	if err := h.Service.CreateTvSeries(&tvSeries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Tv series created",
	})
}

func atoiSafe(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}
