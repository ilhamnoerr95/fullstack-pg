package domain

import "time"

type PaymentStatus string

const (
	StatusCompleted  PaymentStatus = "success"
	StatusProcessing PaymentStatus = "pending"
	StatusFailed     PaymentStatus = "failed"
)
type Payment struct {
	ID           string    `json:"id"`
	MerchantName string    `json:"merchant_name"`
	Date         time.Time `json:"date"`
	Amount       float64   `json:"amount"`
	Status       PaymentStatus `json:"status"`
}