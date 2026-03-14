package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {

		db, err = sql.Open("postgres", connStr)

		if err == nil {

			err = db.Ping()

			if err == nil {
				log.Println("Connected to PostgreSQL")
				return db, nil
			}
		}

		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	return nil, err
}
