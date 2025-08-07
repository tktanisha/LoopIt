package models

import "time"

type ReturnRequest struct {
	ID          int       `json:"id"`
	OrderID     int       `json:"order_id"`
	RequestedBy int       `json:"requested_by"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
