package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

type apiConfig struct {
}

func main() {
	dbName := "file:./local.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
