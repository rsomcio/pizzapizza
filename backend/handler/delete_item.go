package handler

import (
    "fmt"

    "github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/rsomcio/pizzapizza/models"
	"github.com/rsomcio/pizzapizza/database"
)
// DeleteItem
func DeleteItem(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
        body := new(models.Item)
        if err := c.BodyParser(&body); err != nil {
           	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error parsing dns request", "data": err})
        }
        err := database.DeleteItem(db)
        if err != nil {
            log.Info("error")
            c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error Inserting to DB", "data": err})
        }

	   return c.JSON(fiber.Map{
		  "status": "success",
				"message": "Deleted",
				"data": fmt.Sprintf("%s %f", body.Name, body.Price)})
	}
}
