package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aikidoaikido115/test-go-crud/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/products", controllers.GetProducts)
	api.Get("/products/:id", controllers.GetProductByID)
	api.Post("/products", controllers.CreateProduct)
	api.Put("/products/:id", controllers.UpdateProduct)
	api.Delete("/products/:id", controllers.DeleteProduct)
}
