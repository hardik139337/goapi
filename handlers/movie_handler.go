package handlers

import (
	"api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(movieUsecase *service.MovieService) *MovieHandler {
	return &MovieHandler{movieUsecase}
}

func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	movieID := c.Param("movieID")
	u, err := strconv.ParseUint(movieID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	m, err := h.movieService.GetMovieByID(int(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, m)

}
