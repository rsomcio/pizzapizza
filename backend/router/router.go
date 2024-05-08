package router

import (
    "github.com/jmoiron/sqlx"
    "github.com/gofiber/fiber/v2"

    "github.com/rsomcio/pizzapizza/handler"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
    api := app.Group("/api")
    api.Get("/", handler.Hello())

    v1 := api.Group("/v1")
    v1.Post("/add", handler.CreateItem(db))
    v1.Post("/get", handler.GetItem(db))
    v1.Post("/update", handler.UpdateItem(db))
    v1.Post("/delete", handler.DeleteItem(db))
}
