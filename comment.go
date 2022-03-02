package core

import (
	"context"
	"time"
)

type Comment struct {
	ID              int       `json:"-"`
	MovieID         int       `json:"-"`
	Message         string    `json:"message"`
	AuthorIPAddress string    `json:"author_ip_address"`
	CreatedAt       time.Time `json:"created_at"`
}

type CommentRepository interface {
	Add(ctx context.Context, comment *Comment) error
	Get(ctx context.Context, movieID int) ([]*Comment, error)
}
