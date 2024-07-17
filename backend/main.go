package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/am1macdonald/torontoWild/internal/database"
	"github.com/jackc/pgx/v5"
)

type apiConfig struct {
	db *database.Queries
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
	mux.HandleFunc("GET /api/v1/user/{id}", cfg.HandleGetUser)
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
