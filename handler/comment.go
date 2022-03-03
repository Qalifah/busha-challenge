package handler

import (
	core "busha-challenge"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrExceededCommentMax = errors.New("exceeded comment length max")
	ErrInvalidMovieID     = errors.New("invalid movie_id")
)

func (h *Handler) AddComment(c *gin.Context) {
	req := struct {
		Message string `json:"message"`
	}{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.WithError(err).Error(ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrInvalidRequestBody))
		return
	}

	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		log.WithError(err).Error(ErrInvalidMovieID)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrInvalidMovieID))
		return
	}

	if len(req.Message) > 500 {
		log.Error(ErrExceededCommentMax)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrExceededCommentMax))
		return
	}

	authorIPAddress := c.Request.RemoteAddr
	comment := &core.Comment{
		MovieID:         movieID,
		Message:         req.Message,
		AuthorIPAddress: authorIPAddress,
	}
	err = h.commentRepository.Add(c, comment)
	if err != nil {
		log.WithError(err).Error(ErrUnExpectedError)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrUnExpectedError))
		return
	}

	err = h.movieRepository.IncreaseCommentCount(c, movieID)
	if err != nil {
		log.WithError(err).Error(ErrUnExpectedError)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusCreated, NewResponse("comment successfully created", nil))
}

func (h *Handler) GetComments(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		log.WithError(err).Error(ErrInvalidMovieID)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrInvalidMovieID))
		return
	}

	comments, err := h.commentRepository.Get(c, movieID)
	if err != nil {
		log.WithError(err).Error(ErrUnExpectedError)
		c.JSON(http.StatusBadRequest, NewErrorResponse(ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusOK, NewResponse("", comments))
}
