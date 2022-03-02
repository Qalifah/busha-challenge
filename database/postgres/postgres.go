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
	const format = "postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%d"
	uri := fmt.Sprintf(format, config.Username, config.Password, config.Host, config.Port, config.Database, config.MaxConn)
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

func deleteAllComments(ctx context.Context, conn *pgx.Conn, tableName string) {
	conn.Exec(ctx, fmt.Sprintf("DELETE FROM %v", tableName))
}
