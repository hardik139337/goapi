// models/movie.go
package models

import "time"

type Movie struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	PosterPath  string    `json:"poster_path"`
	Language    string    `json:"language"`
	Overview    string    `json:"overview"`
	ReleaseDate time.Time `json:"release_date"`
}
