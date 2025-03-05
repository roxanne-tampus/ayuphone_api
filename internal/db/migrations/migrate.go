package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	ctx := context.Background()
	// Define command-line flags
	operation := flag.String("operation", "migrate", "Operation to perform: up or down")
	flag.Parse()

	// Load environment variables from .env file
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Load environment variables
	databaseURL := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	if databaseURL == "" || authToken == "" {
		log.Fatalf("TURSO_DATABASE_URL or TURSO_AUTH_TOKEN not set")
	}

	// Set the auth token as a connection parameter
	dsn := fmt.Sprintf("%s?auth_token=%s", databaseURL, authToken)
	sqldb, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer sqldb.Close()

	// Create a Bun DB instance
	db := bun.NewDB(sqldb, sqlitedialect.New())

	// Print all queries to stdout
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	schemaUpPath := "internal/db/migrations/001_create_initial_schema_up.sql"
	schemaDownPath := "internal/db/migrations/001_create_initial_schema_down.sql"
	seederPath := "internal/db/migrations/seeder/seeder.sql"

	switch *operation {
	case "up":
		// Run the schema migration
		if err := runSQLFile(ctx, db, schemaUpPath); err != nil {
			log.Fatalf("failed to run schema migration: %v", err)
		}

		// Run the seeder
		if err := runSQLFile(ctx, db, seederPath); err != nil {
			log.Fatalf("failed to run seeder: %v", err)
		}

		fmt.Println("Migration and seeding completed successfully")

	case "down":
		// Run the schema rollback
		if err := runSQLFile(ctx, db, schemaDownPath); err != nil {
			log.Fatalf("failed to run schema rollback: %v", err)
		}

		fmt.Println("Rollback completed successfully")

	default:
		log.Fatalf("unknown operation: %s", *operation)
	}
}

func runSQLFile(ctx context.Context, db *bun.DB, filePath string) error {
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	_, err = db.ExecContext(ctx, string(sqlBytes))
	if err != nil {
		return fmt.Errorf("failed to execute SQL file %s: %w", filePath, err)
	}

	return nil
}
