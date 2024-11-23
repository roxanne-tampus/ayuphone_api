# Ayuphone API

Ayuphone API is a backend service for managing device repairs and transactions. It provides endpoints for creating, updating, and retrieving transactions, as well as managing users, devices, and issues.

## Project Overview

This project is built using Go and BunDB for database interactions. It uses SQLite for development and Turso for production. The API is designed to handle various roles such as superadmin, admin, customer, and technician.

## Prerequisites

- Go 1.16 or higher
- SQLite3
- Turso database for production

## Running Migrations

To run the migrations and seed the database, use the following commands:

# Local Database

```sh
make migrate-up
make migrate-down
```

# Live Database

```sh
make migrate-prod-up
make migrate-prod-down
```

## Running the Application

```sh
make run
```
