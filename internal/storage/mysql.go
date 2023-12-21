package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	db, _ := sql.Open("mysql", "dellis:@/shud")

	return db, nil
}
