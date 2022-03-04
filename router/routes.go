package router

import (
	"busha-challenge/handler"
	"github.com/gin-gonic/gin"
)

func Setup(handler *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/movie", handler.GetMovies)
	r.GET("/:movie_id/comments", handler.GetComments)
	r.POST("/:movie_id/comment", handler.AddComment)
	r.GET("/:movie_id/character", handler.GetCharacters)

	return r
}
