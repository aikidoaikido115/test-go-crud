package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aikidoaikido115/test-go-crud/storage"
	"github.com/aikidoaikido115/test-go-crud/models"
	"fmt"
)

func GetProducts(c *fiber.Ctx) error {
	return c.JSON(storage.Products)
}

func GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	for _, product := range storage.Products {
		if product.ID == id {
			return c.JSON(product)
		}
	}

	return c.Status(404).SendString("Product not found")
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	newProduct.ID = len(storage.Products) + 1
	storage.Products = append(storage.Products, newProduct)
	return c.Status(201).JSON(newProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var updateProduct models.Product
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}	

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	for i, product := range storage.Products {
		if product.ID == id {
			storage.Products[i].Name = updateProduct.Name
			storage.Products[i].Price = updateProduct.Price
			storage.Products[i].CategoryID = updateProduct.CategoryID
			return c.SendString("Product updated")
		}
	}

	return c.Status(404).SendString("Product not found")

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	for i, product := range storage.Products {
		if product.ID == id {
			storage.Products = append(storage.Products[:i], storage.Products[i+1:]...)
			return c.SendString("Product deleted")
		}
	}

	return c.Status(404).SendString("Product not found")
}


//////////////////////////////////////////////////////

func GetProductWithCategory(c *fiber.Ctx) error {
	category := c.Query("category", "null")
	var ProductWithCategory []models.Product
	category_id := -1

	fmt.Println((category))
	
	for _, cate := range storage.Categories {
		if cate.Name == category {
			category_id = cate.ID
		}
	}

	for _, product := range storage.Products {
		if product.CategoryID == category_id {
			ProductWithCategory = append(ProductWithCategory, product)
		}
	}


	if ProductWithCategory == nil {
		return c.Status(404).SendString("Product not found") 
	}

	return c.JSON(ProductWithCategory)
}