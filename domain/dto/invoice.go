package dto

import "time"

type InvoiceRequest struct {
	CustomerID int       `json:"customer_id" validate:"required,gt=0"`
	Amount     float64   `json:"amount" validate:"required,gt=0"`
	Currency   string    `json:"currency" validate:"required,oneof=IDR USD"`
	DueDate    time.Time `json:"due_date" validate:"required"`
}

type InvoiceResponse struct {
	ID         int        `json:"id"`
	CustomerID int        `json:"customer_id"`
	Amount     float64    `json:"amount"`
	PaidAmount float64    `json:"paid_amount"`
	Currency   string     `json:"currency"`
	DueDate    time.Time  `json:"due_date"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
