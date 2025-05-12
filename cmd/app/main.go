package main

import (
	"context"
	"database/sql"
	transport "github.com/dhnikolas/state-manager/internal/http"
	_ "github.com/dhnikolas/state-manager/internal/migrations"
	"github.com/dhnikolas/state-manager/internal/repositories/resources"
	"github.com/dhnikolas/state-manager/internal/services/states"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL environment variable not set")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
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

	resourceRepository := resources.NewResourceRepository(pool)
	stateService := states.NewStateService(resourceRepository)
	currentRouter := transport.NewRouter(stateService)

	log.Fatal(http.ListenAndServe(":8080", currentRouter))
}
