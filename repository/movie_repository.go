package repository

import (
	"api/models"
	"time"

	"gorm.io/gorm"
)

var cacheExpiration = time.Minute

type MovieRepositoryI interface {
	GetMovieByID(id uint) (*models.Movie, error)
}

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db}
}

func (r *MovieRepository) GetMovieByID(id int) (*models.Movie, error) {

	var movie models.Movie
	err := r.db.First(&movie, id).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
