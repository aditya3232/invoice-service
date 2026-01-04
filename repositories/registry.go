package repositories

import (
	invoiceRepo "invoice-service/repositories/invoice"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetInvoice() invoiceRepo.IInvoiceRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetInvoice() invoiceRepo.IInvoiceRepository {
	return invoiceRepo.NewInvoiceRepository(r.db)
}
