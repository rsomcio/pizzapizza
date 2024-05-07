package main


import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/rsomcio/pizzapizza/server"
	"github.com/rsomcio/pizzapizza/database"
)

func main() {
	fmt.Println("PizzaPizza")
	db, _ := database.NewDatabase()
	app, _ := server.NewServer(db)
	app.SetupRoutes()

	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
