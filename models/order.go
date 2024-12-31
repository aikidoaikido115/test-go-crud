package models

type Order struct {
	ID        int   `json:"id"`
	ProductID int   `json:"product_id"`
	UserID    int   `json:"user_id"`
	Quantity  int   `json:"quantity"`
}