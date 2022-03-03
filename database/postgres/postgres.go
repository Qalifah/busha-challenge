package postgres

import (
	"busha-challenge/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func New(ctx context.Context, config *config.PostgresConfig) (*pgx.Conn, error) {
	const format = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	uri := fmt.Sprintf(format, config.Username, config.Password, config.Host, config.Port, config.Database)
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createTestDB(ctx context.Context) (*pgx.Conn, error) {
	testUri := os.Getenv("TEST_DB_URI")
	return pgx.Connect(ctx, testUri)
}
