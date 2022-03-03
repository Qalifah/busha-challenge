package main

import (
	"busha-challenge/database/postgres"
	"busha-challenge/database/redis"
	"busha-challenge/handler"
	"busha-challenge/router"
	"context"
	"flag"
	"os"

	"busha-challenge/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var configPath *string

func init() {
	configPath = flag.String("config_path", "", "path to config file")
	flag.Parse()
	if configPath == nil {
		log.Fatal("-config_path flag is required")
	}
}

func main() {
	file, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("unable to open config file: %v", err)
	}

	cfg := &config.BaseConfig{}
	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	ctx := context.Background()
	conn, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}
	defer conn.Close(ctx)

	commentRepository := postgres.NewCommentRepository(conn)
	characterRepository := postgres.NewCharacterRepository(conn)

	redisClient, err := redis.New(ctx, cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		log.Fatalf("failed to create redis client: %v", err)
	}
	defer redisClient.Close()

	movieRepository := redis.NewMovieRepository(redisClient)

	ctrl := handler.New(movieRepository, commentRepository, characterRepository)
	r := router.Setup(ctrl)

	log.Fatal(r.Run(cfg.ServePort))
}
