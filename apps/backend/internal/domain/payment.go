package domain

import "time"


type Payment struct {
	ID           string    `json:"id"`
	MerchantName string    `json:"merchant_name"`
	Date         time.Time `json:"date"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"`
}