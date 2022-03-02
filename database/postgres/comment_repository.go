package postgres

import (
	core "busha-challenge"
	"context"
	"github.com/jackc/pgx/v4"
)

type commentRepository struct {
	client *pgx.Conn
}

func NewCommentRepository(client *pgx.Conn) *commentRepository {
	return &commentRepository{client: client}
}

func (r *commentRepository) Add(ctx context.Context, comment *core.Comment) error {
	_, err := r.client.Exec(ctx, "INSERT INTO comments (movie_id, message, author_ip_address) VALUES ($1, $2, $3)", comment.MovieID, comment.Message, comment.AuthorIPAddress)
	return err
}

func (r *commentRepository) Get(ctx context.Context, movieID int) ([]*core.Comment, error) {
	var comments []*core.Comment
	rows, err := r.client.Query(ctx, "SELECT * FROM comments WHERE movie_id = $1 ORDER BY created_at DESC", movieID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment core.Comment
		err = rows.Scan(&comment.ID, &comment.MovieID, &comment.Message, &comment.AuthorIPAddress, &comment.CreatedAt)
		comments = append(comments, &comment)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return comments, nil
}
