package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// OpenConnection to a mysql connection
func OpenConnection(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DNS())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	return db, nil
}
