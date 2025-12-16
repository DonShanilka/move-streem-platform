package Handler

import (
	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Service"
)

type EpisodeController struct {
	Service *services.EpisodeService
}

func NewEpisodeController(service *services.EpisodeService) *EpisodeController {
	return &EpisodeController{Service: service}
}

func (c *EpisodeController) CreateEpisode(episode Models.Episode) error {
	return c.Service.SaveEpisode(episode)
}
