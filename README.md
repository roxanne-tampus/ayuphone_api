# Ayuphone API

Ayuphone API is a backend service for managing device repairs and transactions. It provides endpoints for creating, updating, and retrieving transactions, as well as managing users, devices, and issues. This project is built using Go and BunDB for database interactions. It uses SQLite for development and Turso for production. The API is designed to handle various roles such as superadmin, admin, customer, and technician.

## Prerequisites

- Go 1.16 or higher
- SQLite3
- Turso database for production
- WSL2 (Windows Subsystem for Linux) for Windows users

## Installing WSL2 on Windows

To run `make` commands on Windows, you need to install WSL2. Follow these steps:

1. Open PowerShell as Administrator and run:
   ```sh
   wsl --install
   ```
2. Restart your computer if prompted.
3. Set WSL2 as the default version:
   ```sh
   wsl --set-default-version 2
   ```
4. Install a Linux distribution from the Microsoft Store (e.g., Ubuntu).

## Running Migrations

To run the migrations and seed the database, use the following commands:

### Local Database

```sh
make migrate-up
make migrate-down
```

### Live Database

```sh
make migrate-prod-up
make migrate-prod-down
```

## Running the Application

```sh
make run
```
