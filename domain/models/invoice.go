package models

import (
	"invoice-service/constants"
	"time"
)

type Invoice struct {
	ID         int                           `gorm:"primaryKey;autoIncrement"`
	CustomerID int                           `gorm:"not null;index"`
	Amount     float64                       `gorm:"type:numeric(15,2);not null"`
	PaidAmount float64                       `gorm:"type:numeric(15,2);default:0"`
	Currency   string                        `gorm:"type:varchar(10);not null"`
	DueDate    time.Time                     `gorm:"not null"`
	Status     constants.InvoiceStatusString `gorm:"type:varchar(20);not null"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
