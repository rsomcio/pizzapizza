package database

import (
    "fmt"
    "github.com/jmoiron/sqlx"
   	"github.com/gofiber/fiber/v2/log"

)

func GetItem(db *sqlx.DB) error {
    item := `SELECT * FROM item`
    result, err := db.Exec(item)
    if err != nil {
        log.Error("error %v", err)
        return fmt.Errorf("error %v", err)
    }
    log.Info("result %v", result)
    return nil
}
