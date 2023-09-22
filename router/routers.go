package router

import (
	"api/handlers"
	"api/models"
	"api/repository"
	"api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Movie{})

	db.Create(&models.Movie{Title: "iron man", ID: 1})

	movieRepo := repository.NewMovieRepository(db)
	movieUsecase := service.NewMovieService(movieRepo)
	movieHandler := handlers.NewMovieHandler(movieUsecase)

	v1 := r.Group("/v1")
	{
		v1.GET("/movie/:movieID", movieHandler.GetMovieByID)
	}

	return r
}
