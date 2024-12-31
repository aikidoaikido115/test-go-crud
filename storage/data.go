package storage

import "github.com/aikidoaikido115/test-go-crud/models"


var Products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 1500.0, Stock: 999, CategoryID: 1},
	{ID: 2, Name: "Smartphone", Price: 800.0, Stock: 50, CategoryID: 2},
	{ID: 3, Name: "Tablet", Price: 300.0, Stock: 6, CategoryID: 2},
	{ID: 4, Name: "Fake Iphone", Price: 310.0, Stock: 222, CategoryID: 3},
}


var Categories = []models.Category{
	{ID: 1, Name: "Electronics"},
	{ID: 2, Name: "Gadgets"},
	{ID: 3, Name: "Toys"},
	{ID: 4, Name: "Weapons"},
	{ID: 5, Name: "Foods"},
}

var Users = []models.User{
	{ID: 1, Username: "sarai", Email: "sarai@example.com"},
	{ID: 2, Username: "ck", Email: "ck@example.com"},
	{ID: 3, Username: "ck2", Email: "ck2@example.com"},
}

var Orders = []models.Order{
	{ID: 1, ProductID: 1, UserID: 1, Quantity: 2},
	{ID: 1, ProductID: 3, UserID: 1, Quantity: 2},
	{ID: 2, ProductID: 2, UserID: 2, Quantity: 1},
}
