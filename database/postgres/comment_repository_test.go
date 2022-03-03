package postgres

import (
	core "busha-challenge"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCommentRepository_Add(t *testing.T) {
	data := []*core.Comment{
		{1, 3, "Spot On!", "", time.Now()},
		{2, 3, "Thank you!", "", time.Now()},
	}

	ctx := context.Background()
	conn, err := createTestDB(ctx)
	require.NoError(t, err)
	defer conn.Close(ctx)
	commentRepo := NewCommentRepository(conn)

	for _, i := range data {
		require.NoError(t, commentRepo.Add(ctx, i))
	}
}

func TestCommentRepository_Get(t *testing.T) {
	data := []*core.Comment{
		{1, 3, "Spot On!", "", time.Now()},
		{2, 3, "Thank you!", "", time.Now()},
	}

	ctx := context.Background()
	conn, err := createTestDB(ctx)
	require.NoError(t, err)
	defer conn.Close(ctx)
	commentRepo := NewCommentRepository(conn)

	for _, i := range data {
		require.NoError(t, commentRepo.Add(ctx, i))
	}

	_, err = commentRepo.Get(ctx, 3)
	require.NoError(t, err)
}
