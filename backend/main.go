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

func connectDB(ctx context.Context) (*pgx.Conn, *database.Queries) {
	dbUrl := os.Getenv("DB_URL")
	fmt.Println(dbUrl)
	conn, err := pgx.Connect(
		ctx,
		dbUrl,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	return conn, database.New(conn)
}

func connectCache() *valkey.Client {

	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{"valkey:6379"},
	})

	if err != nil {
		panic(err)
	}
	log.Println("Connected to cache")
	return &client
}

func main() {
	ctx := context.Background()
	conn, Queries := connectDB(ctx)
	defer conn.Close(ctx)

	client := connectCache()

	cfg := apiConfig{
		db:     Queries,
		valKey: client,
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
