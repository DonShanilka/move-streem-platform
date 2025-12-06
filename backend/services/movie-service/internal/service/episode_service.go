package services

import (
	"database/sql"
	
	"github.com/DonShanilka/movie-service/internal/models"
	"github.com/DonShanilka/movie-service/internal/repository"
)

type EpisodeService struct {
	Repo *repository.EpisodeRepository
}

func NewEpisodeService(db *sql.DB) *EpisodeService {
	return &EpisodeService{
		Repo: &repository.EpisodeRepository{DB: db},
	}
}

func (s *EpisodeService) SaveEpisode(episode models.Episode) error {
	return s.Repo.SaveEpisode(episode)
}