package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Client struct {
	DB *bun.DB
}

// InitDB initializes the database connection.
func InitDB() (*Client, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Create a new SQL database connection
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if sqldb == nil {
		return nil, fmt.Errorf("failed to create SQL database connection")
	}

	// Initialize the global DB variable
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false), // Set to true to enable query logging
		bundebug.FromEnv(""),
	))

	// Ping the database to ensure connection
	if err := db.PingContext(context.Background()); err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}

	return &Client{DB: db}, nil
}
