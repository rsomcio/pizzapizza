package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"

	"github.com/rsomcio/pizzapizza/database"
	"github.com/rsomcio/pizzapizza/models"
)

// Add
func CreateItem(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(models.Item)
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error parsing dns request", "data": err})
		}
		log.Info("post\n", body.Name)
		err := database.CreateItem(db, &models.Item{
			Name:        body.Name,
			Price:       body.Price,
			Description: body.Description,
			Vendor:      body.Vendor,
			Count:       body.Count,
		})
		if err != nil {
			log.Info("error")
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error Inserting to DB", "data": err})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Hello I'm a Post!",
			"data":    fmt.Sprintf("%s %f", body.Name, body.Price)})
	}
}
