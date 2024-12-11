package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type Client struct {
	DB *bun.DB
}

func NewSQLiteDBClient() (Client, error) {
	mode := os.Getenv("MODE")
	var sqldb *sql.DB
	var err error

	if mode == "dev" {
		dbPath := os.Getenv("SQLITE_DB_PATH")
		if dbPath == "" {
			dbPath = "internal/db/app.db"
		}

		sqldb, err = sql.Open(sqliteshim.ShimName, dbPath)
		if err != nil {
			return Client{}, err
		}

	} else {
		databaseURL := os.Getenv("TURSO_DATABASE_URL")
		authToken := os.Getenv("TURSO_AUTH_TOKEN")

		if databaseURL == "" || authToken == "" {
			return Client{}, fmt.Errorf("TURSO_DATABASE_URL or TURSO_AUTH_TOKEN not set")
		}

		// Set the auth token as a connection parameter
		dsn := fmt.Sprintf("%s?authToken=%s", databaseURL, authToken)
		sqldb, err = sql.Open("libsql", dsn)
		if err != nil {
			return Client{}, err
		}
	}

	// Verify the connection
	if err = sqldb.Ping(); err != nil {
		return Client{}, err
	}

	// Create a Bun DB instance
	db := bun.NewDB(sqldb, sqlitedialect.New())

	// Print all queries to stdout
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return Client{DB: db}, nil
}
