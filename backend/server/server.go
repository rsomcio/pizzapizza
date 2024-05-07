package server

import (
	"github.com/gofiber/fiber/v2"
)
type Server struct {
	web *fiber.App
}

func NewServer() (*Server, error) {
	app := fiber.New()
	app.Get("/", hello)

	return &Server{
		web: app,
	}, nil
}

func (s *Server) Start() error {
	return s.web.Listen(":3000")
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
