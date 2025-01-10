package entities

import (
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusCanceled  PaymentStatus = "CANCELED"
)

type Transaction struct {
	Id int `json:"id"`
	OrderId int `json:"order_id"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	PaymentDate time.Time `json:"payment_date"`
	TotalAmount float64 `json:"total_amount"`
}