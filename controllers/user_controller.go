package controllers

import (
	"github.com/aikidoaikido115/test-go-crud/models"
	"github.com/aikidoaikido115/test-go-crud/storage"
	"github.com/gofiber/fiber/v2"
)

func GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	var userFound *models.User
	for _, user := range storage.Users {
		if user.Username == username {
			userFound = &user
			break
		}
	}

	if userFound == nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   "User not found",
			"message": "The user with the specified username does not exist.",
		})
	}

	return c.JSON(userFound)
}


func SendOrder(c *fiber.Ctx) error {
    var newOrder models.Order

	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

    for _, product := range storage.Products {
		if product.ID == newOrder.ProductID && product.Stock <= 0{
            return c.Status(409).JSON(fiber.Map{
                "error":   "Out of stock",
                "message": "The product is currently out of stock.",
            })
		}
	}
	newOrder.ID = len(storage.Orders) + 1
	storage.Orders = append(storage.Orders, newOrder)

    for i, product := range storage.Products {
		if product.ID == newOrder.ProductID {
            // ลดจำนวน stock ลงไป
			storage.Products[i].Stock -= newOrder.Quantity
		}
    }

    return c.Status(201).JSON(newOrder)
}
