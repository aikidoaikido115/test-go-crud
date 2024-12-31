package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aikidoaikido115/test-go-crud/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/products", controllers.GetProducts)
	api.Get("/products/:id", controllers.GetProductByID)
	api.Get("/productsWithcategory", controllers.GetProductWithCategory)

	api.Post("/products", controllers.CreateProduct)
	api.Put("/products/:id", controllers.UpdateProduct)
	api.Delete("/products/:id", controllers.DeleteProduct)

	///////////////////////////////////
	api.Get("/order", controllers.GetOrder)
	api.Get("/order/:id", controllers.GetOrderDetails)
	api.Get("/orderByUserID/:user_id", controllers.GetOrdersByUserID)
	api.Delete("/order/:id", controllers.DeleteOrder)
	//////////////////////////////
	api.Get("/user/:username", controllers.GetUserByUsername)
	api.Post("/sendOrder", controllers.SendOrder)
}
