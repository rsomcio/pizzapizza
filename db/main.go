package main

import (
    "fmt"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var schema = `
CREATE TABLE IF NOT EXISTS Item (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    vendor VARCHAR(255),
    count INT
);
`

var drop = `
DROP TABLE IF EXISTS Item;
`
func main() {
    connectionString := "postgres://user:pass@localhost:5432/inventorydb?sslmode=disable"
    db, err := sqlx.Open("postgres", connectionString)
    db.SetMaxIdleConns(2)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    if err:= db.Ping(); err != nil {
        panic(err)
    }
    res := db.MustExec(schema)



    fmt.Println("Created table database")
}
