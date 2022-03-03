package core

import (
	"context"
	"time"
)

type Movie struct {
	ID           int       `json:"-"`
	Name         string    `json:"title"`
	OpeningCrawl string    `json:"opening_crawl"`
	CommentCount int       `json:"comment_count"`
	ReleaseDate  time.Time `json:"release_date"`
}

type MovieRepository interface {
	AddMany(ctx context.Context, movies []*Movie) error
	GetAll(ctx context.Context) ([]*Movie, error)
	IncreaseCommentCount(ctx context.Context, movieID int) error
}
