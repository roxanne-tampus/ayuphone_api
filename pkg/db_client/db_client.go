package db_client

import (
	"database/sql"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/ayuphone_api/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Client struct {
	DB       *bun.DB
	DsClient *datastore.Client
}

type Credentials struct {
	Token             string `datastore:"-"`
	HashedCredentials string
}

type PlaidAccess struct {
	Token            string `datastore:"-"`
	PlaidAccessToken string
}

func NewDBClient() (*Client, error) {
	// Define the Data Source Name (DSN)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.Variables.User, config.Variables.Password, config.Variables.DBName, config.Variables.DBHost, config.Variables.DBPort)

	// Parse the configuration for pgx
	pgxConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse DSN: %v", err)
	}

	// Register the configuration and create a database connection pool
	dbURI := stdlib.RegisterConnConfig(pgxConfig)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %v", err)
	}

	// Initialize Bun DB with the connection pool
	db := bun.NewDB(dbPool, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(true), // Enable query logging
	))

	// Initialize Datastore client (optional, depending on your environment)
	// ctx := context.Background()
	// dsClient, err := datastore.NewClient(ctx, config.Variables.ProjectID)
	// if err != nil {
	// 	log.Println("Error setting up Datastore client")
	// 	return nil, err
	// }

	// Return a new Client instance
	return &Client{DB: db}, nil
}
