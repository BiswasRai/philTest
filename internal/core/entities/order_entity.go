package entities

import (
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "PENDING"
    OrderStatusCompleted OrderStatus = "COMPLETED"
    OrderStatusCanceled  OrderStatus = "CANCELED"
)

// Order represents an order entity
type Order struct {
    ID         int         `json:"id"`
    CustomerID int         `json:"customer_id"`
    OrderDate  time.Time   `json:"order_date"`
    Status     OrderStatus `json:"status"`
}