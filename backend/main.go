package main
#hello

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/rsomcio/pizzapizza/server"
	"github.com/rsomcio/pizzapizza/database"
)

const (
    banner = `
________  ___  ________  ________  ________
|\   __  \|\  \|\_____  \|\_____  \|\   __  \
\ \  \|\  \ \  \\|___/  /|\|___/  /\ \  \|\  \
\ \   ____\ \  \   /  / /    /  / /\ \   __  \
    \ \  \___|\ \  \ /  /_/__  /  /_/__\ \  \ \  \
    \ \__\    \ \__\\________\\________\ \__\ \__\
    \|__|     \|__|\|_______|\|_______|\|__|\|__|

________  ___  ________  ________  ________
|\   __  \|\  \|\_____  \|\_____  \|\   __  \
\ \  \|\  \ \  \\|___/  /|\|___/  /\ \  \|\  \
\ \   ____\ \  \   /  / /    /  / /\ \   __  \
    \ \  \___|\ \  \ /  /_/__  /  /_/__\ \  \ \  \
    \ \__\    \ \__\\________\\________\ \__\ \__\
    \|__|     \|__|\|_______|\|_______|\|__|\|__|
`
)
func main() {
	db, _ := database.NewDatabase()
	app, _ := server.NewServer(db)
	app.SetupRoutes()

	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
