package main

import (
	"goapi/vinyl-store/internal/api"
	"goapi/vinyl-store/internal/database"
	"log"
)

func main() {
	db, err := database.StartOrCreateDatabase("vinyl_store.duckdb")
	if err != nil {
		log.Fatalf("Falha no banco de dados: %v", err)
	}
	api.StartServer(db)
	defer db.Close()
}
