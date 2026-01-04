package controllers

import (
	invoiceController "invoice-service/controllers/http/invoice"
	"invoice-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetInvoice() invoiceController.IInvoiceController
}

func NewControllerregistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetInvoice() invoiceController.IInvoiceController {
	return invoiceController.NewInvoiceController(r.service)
}
