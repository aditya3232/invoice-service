package repositories

import (
	invoiceRepo "invoice-service/repositories/invoice"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetCustomer() invoiceRepo.IInvoiceRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetCustomer() invoiceRepo.IInvoiceRepository {
	return invoiceRepo.NewInvoiceRepository(r.db)
}
