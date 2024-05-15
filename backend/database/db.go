package database

import (
    "fmt"
    "github.com/gofiber/fiber/v2/log"

   	"github.com/rsomcio/pizzapizza/config"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func NewDatabase() (*sqlx.DB, error) {
    // To-Do Ray Somcio: condense this
    user := config.Config("DB_USER")
    pass := config.Config("DB_PASSWORD")
    dbName := config.Config("DB_NAME")
    dbHost := config.Config("DB_HOST")

    connectionString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, pass, dbHost, dbName)

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
