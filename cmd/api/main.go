package main

import (
	"context"
	"log"

	"rbp/internal/api"
	"rbp/internal/config"
	"rbp/internal/db/sqlc"

	"rbp/pkg/cache"
)

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	dbConn, err := sqlc.NewDB(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	redisClient := cache.NewRedisClient(cfg.Redis)

	router := api.NewRouter(dbConn, redisClient)

	log.Fatal(router.Run(cfg.Server.Port))
}
