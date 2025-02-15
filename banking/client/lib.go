package client

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func conn() (*sql.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfgPostgres := os.Getenv("CONFIG")
	fmt.Println("config----", cfgPostgres)
	db, err := sql.Open("postgres", cfgPostgres)
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
