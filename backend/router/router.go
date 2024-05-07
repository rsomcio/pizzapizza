package router

import (
    "github.com/jmoiron/sqlx"
    "github.com/gofiber/fiber/v2"

    "github.com/rsomcio/pizzapizza/handler"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
    api := app.Group("/api")
    api.Get("/", handler.Hello())
}
