package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) GetCharacters(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		log.WithError(err).Error(ErrInvalidMovieID)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrInvalidMovieID))
		return
	}
	filterArg := c.Query("filter")
	sortArg := c.Query("sort")
	order := c.Query("order")

	characters, err := h.characterRepository.Get(c, movieID, &filterArg, &sortArg, &order)
	if err != nil {
		log.WithError(err).Error(ErrUnExpectedError)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusOK, NewResponse("", characters))
}
