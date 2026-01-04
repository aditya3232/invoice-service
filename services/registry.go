package services

import (
	"invoice-service/clients"
	"invoice-service/repositories"
	invoiceService "invoice-service/services/invoice"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
}

type IServiceRegistry interface {
	GetInvoice() invoiceService.IInvoiceService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry, client clients.IClientRegistry) IServiceRegistry {
	return &Registry{repository: repository, client: client}
}

func (r *Registry) GetInvoice() invoiceService.IInvoiceService {
	return invoiceService.NewInvoiceService(r.repository, r.client)
}
