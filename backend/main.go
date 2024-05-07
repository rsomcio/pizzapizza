package main


import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/rsomcio/pizzapizza/server"
)
func main() {
	fmt.Println("PizzaPizza")

	app, _ := server.NewServer()
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
