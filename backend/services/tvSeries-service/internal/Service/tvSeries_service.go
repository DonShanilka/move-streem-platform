package Service

import (
	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
)

type TvSerriesService struct {
	Repo *Repository.TvSeriesRepository
}

func NewTvSerriesService(repo *Repository.TvSeriesRepository) *TvSerriesService {
	return &TvSerriesService{Repo: repo}
}

func (s *TvSerriesService) CreateTvSeries(tvSerries *Models.Series) error {
	return s.Repo.Create(tvSerries)
}
