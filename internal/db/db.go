package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnMySQL() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	return db
}
