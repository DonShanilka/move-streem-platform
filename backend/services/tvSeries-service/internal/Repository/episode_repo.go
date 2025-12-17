package Repository

import (
	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"gorm.io/gorm"
)

type EpisodeRepository struct {
	DB *gorm.DB
}

func (r *EpisodeRepository) Save(episode *Models.Episode) error {
	return r.DB.Create(episode).Error
}

func (r *EpisodeRepository) Update(episode *Models.Episode) error {
	return r.DB.Save(episode).Error
}
