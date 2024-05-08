package models

import (
    "time"
)

type Item struct {
    ID          int     `db:"id" json:"id"`
    Name        string  `db:"name" json:"name"`
    Description string  `db:"description" json:"description"`
    Price       float64 `db:"price" json:"price"`
    Vendor      string  `db:"vendor" json:"vendor"`
    Count       int     `db:"count" json:"count"`
    CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
    DeletedAt   time.Time `db:"deletedAt" json:"deletedAt"`
    ModifiedAt  time.Time `db:"modifiedAt" json:"modifiedAt"`
}
