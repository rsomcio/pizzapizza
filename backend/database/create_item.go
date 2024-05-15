package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/rsomcio/pizzapizza/models"
)

func CreateItem(db *sqlx.DB, data *models.Item) error {
	item := `INSERT INTO item (name, price, description, vendor, count) VALUES ($1, $2, $3, $4, $5)`
	result, err := db.Exec(item, data.Name, data.Price, data.Description, data.Vendor, data.Count)
	if err != nil {
		log.Error("error %v", err)
		return fmt.Errorf("error %v", err)
	}
	log.Info("result %v", result)
	return nil
}
