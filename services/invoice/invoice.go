package services

import (
	"context"
	"errors"
	"invoice-service/clients"
	"invoice-service/constants"
	errConstant "invoice-service/constants/error"
	"invoice-service/domain/dto"
	"invoice-service/repositories"
	"log"
	"time"
)

type InvoiceService struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
}

type IInvoiceService interface {
	FindByID(context.Context, int) (*dto.InvoiceResponse, error)
	MarkOverdue(context.Context, int) (*dto.InvoiceMarkOverdueResponse, error)
	Create(context.Context, *dto.InvoiceRequest) (*dto.InvoiceResponse, error)
	FindAllWithoutPagination(context.Context, *dto.InvoiceRequestParam) ([]dto.InvoiceResponse, error)
	StartMarkOverdueJob(context.Context, time.Duration) error
	HandlePayment(context.Context, *dto.PaymentData) error
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

func (s *InvoiceService) MarkOverdue(ctx context.Context, invoiceID int) (*dto.InvoiceMarkOverdueResponse, error) {
	invoice, err := s.FindByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	// sudah lunas, tidak boleh overdue
	if invoice.PaidAmount >= invoice.Amount {
		return nil, errConstant.ErrMarkOverdueAlreadyFullPaid
	}

	// belum lewat due date
	if time.Now().Before(invoice.DueDate) {
		return nil, errConstant.ErrMarkOverdueNotOverdueYet
	}

	// invalid status
	if invoice.Status != constants.Unpaid &&
		invoice.Status != constants.PartiallyPaid {
		return nil, errConstant.ErrMarkOverdueInvalidStatus
	}

	updateReq := &dto.InvoiceUpdateRequest{
		Status:    constants.Overdue,
		UpdatedAt: time.Now(),
	}

	err = s.repository.GetInvoice().Update(ctx, updateReq, invoiceID)
	if err != nil {
		return nil, err
	}

	response := dto.InvoiceMarkOverdueResponse{
		ID:     invoiceID,
		Status: constants.Overdue,
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

func (s *InvoiceService) StartMarkOverdueJob(ctx context.Context, interval time.Duration) error {
	jobName := "invoice-overdue-job"
	log.Printf("[%s] STARTED | interval=%s\n", jobName, interval)
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Printf("[%s] RUNNING | checking overdue invoices\n", jobName)

				if err := s.repository.GetInvoice().MarkOverdueInvoices(ctx); err != nil {
					log.Printf("[%s] ERROR | mark overdue failed | err=%v\n", jobName, err)
				} else {
					log.Printf("[%s] SUCCESS | overdue invoices updated\n", jobName)
				}

			case <-ctx.Done():
				ticker.Stop()
				log.Printf("[%s] STOPPED | context cancelled\n", jobName)
				return
			}
		}
	}()

	return nil
}

func (s *InvoiceService) HandlePayment(ctx context.Context, req *dto.PaymentData) error {
	invoice, err := s.FindByID(ctx, req.InvoiceID)
	if err != nil {
		return err
	}

	newPaidAmount := invoice.PaidAmount + req.Amount

	if newPaidAmount > invoice.Amount {
		return errors.New("paid amount exceeds invoice amount")
	}

	var newStatus string
	switch {
	case newPaidAmount == invoice.Amount:
		newStatus = string(constants.Paid)
	case newPaidAmount > 0:
		newStatus = string(constants.PartiallyPaid)
	default:
		newStatus = string(constants.Unpaid)
	}

	updateReq := &dto.InvoiceUpdateRequest{
		Status:     constants.InvoiceStatusString(newStatus),
		PaidAmount: newPaidAmount,
		UpdatedAt:  time.Now(),
	}

	err = s.repository.GetInvoice().Update(ctx, updateReq, req.InvoiceID)
	if err != nil {
		return err
	}

	return nil
}
