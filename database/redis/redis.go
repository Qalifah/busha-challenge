package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func New(ctx context.Context, addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func createRedisTestDB(ctx context.Context) (*redis.Client, error) {
	return New(ctx, "localhost:6379", "", 0)
}
