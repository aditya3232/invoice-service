package dto

import (
	"invoice-service/constants"
	"time"
)

type InvoiceRequest struct {
	CustomerID int                           `json:"customer_id" validate:"required,gt=0"`
	Amount     float64                       `json:"amount" validate:"required,gt=0"`
	Currency   string                        `json:"currency" validate:"required,oneof=IDR USD"`
	DueDate    string                        `json:"due_date" validate:"required,datetime=2006-01-02"`
	Status     constants.InvoiceStatusString `json:"status"`
}

type InvoiceUpdateRequest struct {
	Status    constants.InvoiceStatusString `json:"status"`
	UpdatedAt time.Time                     `json:"updated_at"`
}

type InvoiceRequestParam struct {
	CustomerID *int `form:"customer_id"`
}

type InvoiceResponse struct {
	ID         int                           `json:"id"`
	CustomerID int                           `json:"customer_id"`
	Amount     float64                       `json:"amount"`
	PaidAmount float64                       `json:"paid_amount"`
	Currency   string                        `json:"currency"`
	DueDate    time.Time                     `json:"due_date"`
	Status     constants.InvoiceStatusString `json:"status"`
	CreatedAt  *time.Time                    `json:"created_at"`
	UpdatedAt  *time.Time                    `json:"updated_at"`
}

type InvoiceMarkOverdueResponse struct {
	ID     int                           `json:"id"`
	Status constants.InvoiceStatusString `json:"status"`
}
