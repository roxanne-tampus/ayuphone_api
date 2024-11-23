package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	// Define command-line flags
	operation := flag.String("operation", "migrate", "Operation to perform: up or down")
	flag.Parse()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load environment variables
	databaseURL := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	if databaseURL == "" || authToken == "" {
		log.Fatalf("TURSO_DATABASE_URL or TURSO_AUTH_TOKEN not set")
	}

	// Set the auth token as a connection parameter
	dsn := fmt.Sprintf("%s?auth_token=%s", databaseURL, authToken)
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Determine the schema file to use based on the operation
	var schemaPath string
	if *operation == "down" {
		schemaPath = "database/migrations/schema/001_create_tables_down.sql"
	} else {
		schemaPath = "database/migrations/schema/001_create_tables_up.sql"
	}

	// Read the SQL schema file
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	// Execute the SQL schema
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}
	var operationResponse string
	if *operation == "down" {
		operationResponse = "dropped"
	} else {
		operationResponse = "migrated"
	}
	fmt.Printf("Database schema %s successfully\n", operationResponse)
}
