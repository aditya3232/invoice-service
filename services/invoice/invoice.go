package services

import (
	"context"
	"invoice-service/clients"
	"invoice-service/constants"
	errConstant "invoice-service/constants/error"
	"invoice-service/domain/dto"
	"invoice-service/repositories"
)

type InvoiceService struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
}

type IInvoiceService interface {
	FindByID(context.Context, int) (*dto.InvoiceResponse, error)
	Create(context.Context, *dto.InvoiceRequest) (*dto.InvoiceResponse, error)
	FindAllWithoutPagination(context.Context, *dto.InvoiceRequestParam) ([]dto.InvoiceResponse, error)
}

func NewInvoiceService(repository repositories.IRepositoryRegistry, client clients.IClientRegistry) IInvoiceService {
	return &InvoiceService{repository: repository, client: client}
}

func (s *InvoiceService) FindByID(ctx context.Context, id int) (*dto.InvoiceResponse, error) {
	invoice, err := s.repository.GetInvoice().FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.InvoiceResponse{
		ID:         invoice.ID,
		CustomerID: invoice.CustomerID,
		Amount:     invoice.Amount,
		PaidAmount: invoice.PaidAmount,
		Currency:   invoice.Currency,
		DueDate:    invoice.DueDate,
		Status:     invoice.Status,
		CreatedAt:  invoice.CreatedAt,
		UpdatedAt:  invoice.UpdatedAt,
	}

	return &response, nil
}

func (s *InvoiceService) Create(ctx context.Context, req *dto.InvoiceRequest) (*dto.InvoiceResponse, error) {
	_, err := s.client.GetCustomer().FindByID(ctx, req.CustomerID)
	if err != nil {
		return nil, errConstant.ErrCustomerNotFound
	}

	status := constants.Unpaid
	invoice, err := s.repository.GetInvoice().Create(ctx, &dto.InvoiceRequest{
		CustomerID: req.CustomerID,
		Amount:     req.Amount,
		Currency:   req.Currency,
		DueDate:    req.DueDate,
		Status:     status,
	})

	if err != nil {
		return nil, err
	}

	response := &dto.InvoiceResponse{
		ID:         invoice.ID,
		CustomerID: invoice.CustomerID,
		Amount:     invoice.Amount,
		PaidAmount: invoice.PaidAmount,
		Currency:   invoice.Currency,
		DueDate:    invoice.DueDate,
		Status:     status,
		CreatedAt:  invoice.CreatedAt,
		UpdatedAt:  invoice.UpdatedAt,
	}

	return response, nil
}

func (s *InvoiceService) FindAllWithoutPagination(ctx context.Context, req *dto.InvoiceRequestParam) ([]dto.InvoiceResponse, error) {
	invoices, err := s.repository.GetInvoice().FindAllWithoutPagination(ctx, req)
	if err != nil {
		return nil, err
	}

	invoiceResults := make([]dto.InvoiceResponse, 0, len(invoices))
	for _, invoice := range invoices {
		invoiceResults = append(invoiceResults, dto.InvoiceResponse{
			ID:         invoice.ID,
			CustomerID: invoice.CustomerID,
			Amount:     invoice.Amount,
			PaidAmount: invoice.PaidAmount,
			Currency:   invoice.Currency,
			DueDate:    invoice.DueDate,
			Status:     invoice.Status,
			CreatedAt:  invoice.CreatedAt,
			UpdatedAt:  invoice.UpdatedAt,
		})
	}

	return invoiceResults, nil
}
