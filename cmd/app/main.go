package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/reconcile-kit/state-manager/config"
	transport "github.com/reconcile-kit/state-manager/internal/http"
	_ "github.com/reconcile-kit/state-manager/internal/migrations"
	"github.com/reconcile-kit/state-manager/internal/repositories/events"
	"github.com/reconcile-kit/state-manager/internal/repositories/resources"
	"github.com/reconcile-kit/state-manager/internal/services/states"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

func main() {

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		builderURL, err := config.BuildPostgresDSN()
		if err != nil {
			panic(err)
		}
		dbURL = builderURL
	}

	redisURL, err := config.RedisURL()
	if err != nil {
		panic(err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	skipTLS := os.Getenv("REDIS_SKIP_TLS")

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("invalid REDIS_URL: %v", err)
	}

	if skipTLS == "" {
		opt.TLSConfig = &tls.Config{}
	}

	redisClient := redis.NewClient(opt)
	defer redisClient.Close()
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	dbConn, err := sql.Open("pgx", dbURL)
	if err != nil {
		panic(err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(dbConn, "/var"); err != nil {
		panic(err)
	}

	eventsRepo := events.NewRedisRepository(redisClient)
	resourceRepository := resources.NewResourceRepository(pool)
	stateService := states.NewStateService(resourceRepository, eventsRepo)
	currentRouter := transport.NewRouter(stateService)

	log.Fatal(http.ListenAndServe(":"+serverPort, currentRouter))
}
