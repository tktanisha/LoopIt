package models

import "time"

type BuyingRequest struct {
	ID          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	RequestedBy int       `json:"requested_by"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
