package handler

import core "busha-challenge"

type Handler struct {
	movieRepository     core.MovieRepository
	commentRepository   core.CommentRepository
	characterRepository core.CharacterRepository
}

func New(movieRepository core.MovieRepository, commentRepository core.CommentRepository, charRepository core.CharacterRepository) *Handler {
	return &Handler{movieRepository: movieRepository, commentRepository: commentRepository, characterRepository: charRepository}
}
