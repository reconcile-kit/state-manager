package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/reconcile-kit/state-manager/config"
	"github.com/reconcile-kit/state-manager/internal/auth"
	transport "github.com/reconcile-kit/state-manager/internal/http"
	_ "github.com/reconcile-kit/state-manager/internal/migrations"
	"github.com/reconcile-kit/state-manager/internal/repositories/events"
	"github.com/reconcile-kit/state-manager/internal/repositories/permissions"
	"github.com/reconcile-kit/state-manager/internal/repositories/resources"
	"github.com/reconcile-kit/state-manager/internal/services/states"
	"github.com/redis/go-redis/v9"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Bearer <token>". Required only when AUTH_ENABLED=true.
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
	defer pool.Close()

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
	defer dbConn.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	log.Println("Running database migrations...")
	if err := goose.Up(dbConn, "/var"); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migrations done")

	authCfg, err := config.AuthConfig()
	if err != nil {
		log.Fatalf("invalid auth config: %v", err)
	}
	var (
		authenticator *auth.Authenticator
		authorizer    auth.Authorizer = auth.NoopAuthorizer{}
	)
	if authCfg.Enabled {
		verifier, err := auth.NewVerifier(context.Background(), authCfg)
		if err != nil {
			log.Fatalf("failed to init token verifier: %v", err)
		}
		var ruleStore auth.RuleStore
		if authCfg.PermissionsSource != auth.SourceClaims {
			ruleStore = permissions.NewPermissionsRepository(pool)
		}
		authenticator, err = auth.NewAuthenticator(authCfg, verifier, ruleStore)
		if err != nil {
			log.Fatalf("failed to init authenticator: %v", err)
		}
		authorizer = auth.RulesAuthorizer{}
		log.Printf("Authorization enabled, permissions source: %s", authCfg.PermissionsSource)
	} else {
		log.Println("Authorization disabled")
	}

	eventsRepo := events.NewRedisRepository(redisClient)
	resourceRepository := resources.NewResourceRepository(pool)
	stateService := states.NewStateService(resourceRepository, eventsRepo, authorizer)
	currentRouter := transport.NewRouter(stateService, authenticator)

	log.Fatal(http.ListenAndServe(":"+serverPort, currentRouter))
}
