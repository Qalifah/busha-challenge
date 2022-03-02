package redis

import (
	core "busha-challenge"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type movieRepository struct {
	client *redis.Client
}

func NewMovieRepository(client *redis.Client) *movieRepository {
	return &movieRepository{client: client}
}

func (m *movieRepository) AddMany(ctx context.Context, movies []*core.Movie) error {
	for _, i := range movies {
		err := m.set(ctx, strconv.Itoa(i.ID), i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *movieRepository) GetAll(ctx context.Context) ([]*core.Movie, error) {
	var movies []*core.Movie
	keys := m.client.HKeys(ctx, "movies").Val()
	for _, k := range keys {
		movie, err := m.get(ctx, k)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (m *movieRepository) IncreaseCommentCount(ctx context.Context, movieID string) error {
	movie, err := m.get(ctx, movieID)
	if err != nil {
		return err
	}
	movie.CommentCount++
	return m.set(ctx, movieID, movie)
}

func (m *movieRepository) get(ctx context.Context, movieID string) (*core.Movie, error) {
	var movie core.Movie
	val, err := m.client.HGet(ctx, "movies", movieID).Result()
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(val), &movie)
	return &movie, nil
}

func (m *movieRepository) set(ctx context.Context, movieID string, movie *core.Movie) error {
	newValue, err := json.Marshal(movie)
	if err != nil {
		return err
	}

	err = m.client.HSet(ctx, "movies", movieID, newValue).Err()
	return err
}
