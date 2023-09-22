package service

import (
	"api/models"
	"api/redis"
	"api/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type MovieServiceI interface {
	GetMovieByID(id uint) (*models.Movie, error)
}
type MovieService struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(movieRepo *repository.MovieRepository) *MovieService {
	return &MovieService{movieRepo}
}

func (u *MovieService) GetMovieByID(id int) (*models.Movie, error) {
	cacheKey := fmt.Sprintf("movie:%d", id)

	cachedMovieJSON, err := redis.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var cachedMovie models.Movie
		if err := json.Unmarshal([]byte(cachedMovieJSON), &cachedMovie); err != nil {
			return nil, err
		}
		return &cachedMovie, nil
	}

	movie, err := u.movieRepo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}

	movieJSON, err := json.Marshal(movie)
	if err != nil {
		return nil, err
	}
	err = redis.RedisClient.Set(context.Background(), cacheKey, movieJSON, cacheExpiration).Err()

	if err != nil {
		return nil, err
	}

	return movie, err
}

var cacheExpiration = time.Minute
