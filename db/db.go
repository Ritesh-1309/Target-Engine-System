package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening DB: %w", err)
	}
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error pinging DB: %w", err)
	}
	return nil
}
