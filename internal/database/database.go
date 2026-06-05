package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/duckdb/duckdb-go/v2"
)

const DATASOURCE_PATH = "../internal/database/data/vinyl_store.duckdb"
const DB_DRIVER_NAME = "duckdb"

func StartOrCreateDatabase(dataSourcePath string) (*sql.DB, error) {
	db, dbErr := sql.Open(DB_DRIVER_NAME, dataSourcePath)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	if err := migrate(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("Erro ao aplicar migrações: %w", err)
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	query := `
	CREATE SEQUENCE IF NOT EXISTS album_id_sequence;
	CREATE TABLE IF NOT EXISTS albums (
		id INTEGER PRIMARY KEY DEFAULT nextval('album_id_sequence'),
		title VARCHAR NOT NULL,
		artist VARCHAR NOT NULL,
		price DECIMAL(10, 2) NOT NULL
	);`

	_, err := db.Exec(query)
	return err
}
