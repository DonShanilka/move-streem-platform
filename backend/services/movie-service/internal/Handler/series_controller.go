package Handler

import (
	"github.com/DonShanilka/movie-service/internal/Models"
	"github.com/DonShanilka/movie-service/internal/Service"
)

type SeriesController struct {
	Service *services.SeriesService
}

func NewSeriesController(service *services.SeriesService) *SeriesController {
	return &SeriesController{Service: service}
}

func (c *SeriesController) CreateSeries(service Models.Series) error {
	return c.Service.SaveSeries(service)
}
