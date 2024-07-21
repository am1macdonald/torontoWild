package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/am1macdonald/torontoWild/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/valkey-io/valkey-go"
)

type apiConfig struct {
	db     *database.Queries
	valKey *valkey.Client
}

func main() {
	ctx := context.Background()
	dbUrl := os.Getenv("DB_URL")
	fmt.Println(dbUrl)
	conn, err := pgx.Connect(
		ctx,
		dbUrl,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	log.Println("Connected to database")

	cfg := apiConfig{
		db: database.New(conn),
	}

	mux := http.NewServeMux()

	// user sessions
	mux.HandleFunc("POST /api/v1/sign-in", cfg.HandleSignIn)
	mux.HandleFunc("POST /api/v1/sign-out", cfg.HandleSignIn)
	mux.HandleFunc("GET /api/v1/magic-link", cfg.HandleSignIn)

	// wildlife sightings
	mux.HandleFunc("POST /api/v1/sightings", cfg.HandleSignIn)
	mux.HandleFunc("GET /api/v1/sightings", cfg.HandleSignIn)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
