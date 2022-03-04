package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

var (
	ErrUnExpectedError = errors.New("an error occurred, please try again")
)

func (h *Handler) GetMovies(c *gin.Context) {
	movies, err := h.movieRepository.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(ErrUnExpectedError))
		return
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ReleaseDate.Before(movies[j].ReleaseDate)
	})

	c.JSON(http.StatusOK, NewResponse("all movies successfully returned", movies))
}
