package redis

import (
	core "busha-challenge"
	"context"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func TestMovieRepository_AddMany(t *testing.T) {
	data := []*core.Movie{
		{1, "Naruto", "In the beginning.....", 0, time.Now()},
		{2, "Demon Slayer", "HA ha ha", 2, time.Now()},
	}
	ctx := context.Background()
	client, err := createRedisTestDB(ctx)
	require.NoError(t, err)
	defer client.Close()

	movieRepo := NewMovieRepository(client)
	err = movieRepo.AddMany(ctx, data)
	require.NoError(t, err)
}

func TestMovieRepository_GetAll(t *testing.T) {
	data := []*core.Movie{
		{1, "Naruto", "In the beginning.....", 0, time.Now()},
		{2, "Demon Slayer", "HA ha ha", 2, time.Now()},
	}
	ctx := context.Background()
	client, err := createRedisTestDB(ctx)
	require.NoError(t, err)
	client.FlushDB(ctx)
	defer client.Close()

	movieRepo := NewMovieRepository(client)
	err = movieRepo.AddMany(ctx, data)
	require.NoError(t, err)

	movies, err := movieRepo.GetAll(ctx)
	require.NoError(t, err)
	require.Equal(t, len(data), len(movies))
}

func TestMovieRepository_IncreaseCommentCount(t *testing.T) {
	movie := &core.Movie{
		ID: 1, Name: "COD",
		OpeningCrawl: "Choose your death!",
		CommentCount: 0,
		ReleaseDate:  time.Now(),
	}
	ctx := context.Background()
	client, err := createRedisTestDB(ctx)
	require.NoError(t, err)
	defer client.Close()

	movieRepo := NewMovieRepository(client)
	movieID := strconv.Itoa(movie.ID)
	err = movieRepo.set(ctx, movieID, movie)
	require.NoError(t, err)

	err = movieRepo.IncreaseCommentCount(ctx, movieID)
	require.NoError(t, err)

	movie, err = movieRepo.get(ctx, movieID)
	require.NoError(t, err)
	require.Equal(t, 1, movie.CommentCount)
}
