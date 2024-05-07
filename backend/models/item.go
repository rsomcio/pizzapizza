package models

type Item struct {
    ID          int     `db:"id" json:"id"`
    Name        string  `db:"name" json:"name"`
    Description string  `db:"description" json:"description"`
    Price       float64 `db:"price" json:"price"`
    Vendor      string  `db:"vendor" json:"vendor"`
    Count       int     `db:"count" json:"count"`
}
