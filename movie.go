package core

import (
	"context"
	"time"
)

type Movie struct {
	ID           int       `json:"-"`
	Name         string    `json:"name"`
	OpeningCrawl string    `json:"opening_crawl"`
	CommentCount int       `json:"comment_count"`
	ReleaseDate  time.Time `json:"-"`
}

type MovieRepository interface {
	AddMany(ctx context.Context, movies []*Movie) error
	GetAll(ctx context.Context) ([]*Movie, error)
	IncreaseCommentCount(ctx context.Context, movieID string) error
}
