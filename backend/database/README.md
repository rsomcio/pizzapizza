package main

import (
"fmt"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // PostgreSQL driver

)

var schema = `CREATE TABLE IF NOT EXISTS Item (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    vendor VARCHAR(255),
    count INT
);`

var drop = `DROP TABLE IF EXISTS Item;`

var insert = `     INSERT INTO Item (name, description, price, vendor, count)
    VALUES
        ('Product A', 'Description of Product A', 10.99, 'Vendor X', 100),
        ('Product B', 'Description of Product B', 15.49, 'Vendor Y', 75),
        ('Product C', 'Description of Product C', 20.99, 'Vendor Z', 50);
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
    fmt.Printf("schema %+v\n", res)

    item := `INSERT INTO item (name, price) VALUES ($1, $2)`
    // res = db.MustExec(insert)
    result, err := db.Exec(item, "Hong Kong", 852.00)
    if err != nil {
        fmt.Errorf("error %v", err)
        panic(err)
    }
    fmt.Printf("result %+v\n", result)


    fmt.Println("Connected to the database")

}
