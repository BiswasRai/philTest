package entities

type OrderItem struct {
	Id int `json:"id"`
	OrderID int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}