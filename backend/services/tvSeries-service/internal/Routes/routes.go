package Routes

import (
	"net/http"

	"github.com/DonShanilka/tvSeries-service/internal/Handler"
)

func RegisterTvSeriesRoutes(mux *http.ServeMux, h *Handler.TvSeriesHandler) {
	mux.HandleFunc("/api/series/createSeries", h.CreateTvSeries) // POST
	//mux.HandleFunc("/api/series/add-season", h.AddSeason)
}

func RegisterEpisodeRoutes(mux *http.ServeMux, h *Handler.EpisodeHandler) {
	// Episode routes
	mux.HandleFunc("/api/episodes/create", h.CreateEpisode)
	mux.HandleFunc("/api/episodes/update", h.UpdateEpisode)
	mux.HandleFunc("/api/episodes/delete", h.DeleteEpisode)
	mux.HandleFunc("/api/episodes/getAllEpisode", h.GetAllEpisodes)
}
