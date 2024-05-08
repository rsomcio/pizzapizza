package database

import (
    "fmt"
    "time"

    "github.com/jmoiron/sqlx"
   	"github.com/gofiber/fiber/v2/log"
)

func UpdateItem(db *sqlx.DB) error {
    currentTime := time.Now()

    item := `UPDATE item
    SET udpateat = $1
    WHERE id = <item_id>;`

    result, err := db.Exec(item, currentTime)
    if err != nil {
        log.Error("error %v", err)
        return fmt.Errorf("error %v", err)
    }
    log.Info("result %v", result)
    return nil
}
