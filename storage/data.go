package storage

import "github.com/aikidoaikido115/test-go-crud/models"


var Products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 1500.0, Stock: 20, CategoryID: 1},
	{ID: 2, Name: "Smartphone", Price: 800.0, Stock: 50, CategoryID: 2},
	{ID: 3, Name: "Tablet", Price: 300.0, Stock: 30, CategoryID: 2},
	{ID: 4, Name: "Fake Iphone", Price: 310.0, Stock: 2, CategoryID: 3},
}
