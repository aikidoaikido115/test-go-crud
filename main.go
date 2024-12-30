package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aikidoaikido115/test-go-crud/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}