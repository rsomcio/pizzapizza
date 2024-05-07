package server

import (
    "github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2"

	"github.com/rsomcio/pizzapizza/router"

)
type Server struct {
	web *fiber.App
	db *sqlx.DB
}

func NewServer(db *sqlx.DB) (*Server, error) {
	app := fiber.New()
	app.Get("/", hello)

	return &Server{
		web: app,
		db: db,
	}, nil
}

func (s *Server) Start() error {
	return s.web.Listen(":3000")
}

func (s *Server) SetupRoutes() {
	router.SetupRoutes(s.web, s.db)
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
