package database

import (
    "fmt"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func NewDatabase() (*sqlx.DB, error) {
    connectionString := "postgres://user:pass@localhost:5432/inventorydb?sslmode=disable"
    db, err := sqlx.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }
    // defer db.Close()

    if err:= db.Ping(); err != nil {
        return nil, err
    }
    fmt.Println("Connected to the database")

    return db, nil
}
