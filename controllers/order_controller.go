package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aikidoaikido115/test-go-crud/storage"
	"github.com/aikidoaikido115/test-go-crud/models"
)


func GetOrder(c *fiber.Ctx) error {
	return c.JSON(storage.Orders)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
        return c.Status(400).SendString("Invalid Order ID")
    }

	for i, order := range storage.Orders {
		if order.ID == id {
			storage.Orders = append(storage.Orders[:i], storage.Orders[i+1:]...)
			return c.SendString("Order deleted")
		}
	}

	return c.Status(404).SendString("Order not found")

}
func GetOrdersByUserID(c *fiber.Ctx) error {

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(400).SendString("Invalid User ID")
	}


	var userOrders []models.Order
	for _, order := range storage.Orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}


	if len(userOrders) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error":   "No orders found",
			"message": "This user has no orders.",
		})
	}

	return c.JSON(userOrders)
}

func GetOrderDetails(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(400).SendString("Invalid Order ID")
    }

    var order models.Order
    for _, o := range storage.Orders {
        if o.ID == id {
            order = o
            break
        }
    }
    if order.ID == 0 {
        return c.Status(404).SendString("Order not found")
    }

    var productName string
    for _, product := range storage.Products {
        if product.ID == order.ProductID {
            productName = product.Name
            break
        }
    }

    var username string
    for _, user := range storage.Users {
        if user.ID == order.UserID {
            username = user.Username
            break
        }
    }

    return c.JSON(fiber.Map{
        "id":        order.ID,
        "product":   productName,
        "user":      username,
        "quantity":  order.Quantity,
    })
}
