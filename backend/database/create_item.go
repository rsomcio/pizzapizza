package database

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/rsomcio/pizzapizza/models"
   	"github.com/gofiber/fiber/v2/log"

)

func CreateItem(db *sqlx.DB, data *models.Item) error {
    item := `INSERT INTO item (name, price) VALUES ($1, $2)`
    result, err := db.Exec(item, data.Name, data.Price)
    if err != nil {
        log.Error("error %v", err)
        return fmt.Errorf("error %v", err)
    }
    log.Info("result %v", result)
    return nil
}
