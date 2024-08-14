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

var DBClient *bun.DB

// InitDB initializes the database connection.
func InitDB() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Create a new SQL database connection
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if sqldb == nil {
		return fmt.Errorf("failed to create SQL database connection")
	}

	// Initialize the global DB variable
	DBClient = bun.NewDB(sqldb, pgdialect.New())
	DBClient.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false), // Set to true to enable query logging
		bundebug.FromEnv(""),
	))

	// Ping the database to ensure connection
	if err := DBClient.PingContext(context.Background()); err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return err
	}

	return nil
}
