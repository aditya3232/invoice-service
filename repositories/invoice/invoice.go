package repositories

import (
	"context"
	"errors"
	"invoice-service/constants"
	"invoice-service/domain/dto"
	"invoice-service/domain/models"
	"time"

	errWrap "invoice-service/common/error"
	errConstant "invoice-service/constants/error"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

type IInvoiceRepository interface {
	FindByID(context.Context, int) (*models.Invoice, error)
	Create(context.Context, *dto.InvoiceRequest) (*models.Invoice, error)
	Update(context.Context, *dto.InvoiceUpdateRequest, int) error
	FindAllWithoutPagination(context.Context, *dto.InvoiceRequestParam) ([]models.Invoice, error)
	MarkOverdueInvoices(context.Context) error
}

func NewInvoiceRepository(db *gorm.DB) IInvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) FindByID(ctx context.Context, id int) (*models.Invoice, error) {
	var invoice models.Invoice

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrInvoiceNotFound
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &invoice, nil
}

func (r *InvoiceRepository) Create(ctx context.Context, req *dto.InvoiceRequest) (*models.Invoice, error) {
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrInternalServerError)
	}

	invoice := models.Invoice{
		CustomerID: req.CustomerID,
		Amount:     req.Amount,
		Currency:   req.Currency,
		DueDate:    dueDate,
		Status:     req.Status,
	}

	err = r.db.WithContext(ctx).Create(&invoice).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &invoice, nil
}

func (r *InvoiceRepository) Update(ctx context.Context, req *dto.InvoiceUpdateRequest, invoiceID int) error {
	err := r.db.WithContext(ctx).Model(&models.Invoice{}).Where("id = ?", invoiceID).Updates(req).Error
	if err != nil {
		errWrap.WrapError(errConstant.ErrSQLError)
	}

	return nil
}

func (r *InvoiceRepository) FindAllWithoutPagination(ctx context.Context, req *dto.InvoiceRequestParam) ([]models.Invoice, error) {
	var invoices []models.Invoice
	query := r.db.WithContext(ctx)

	if req.CustomerID != nil {
		query = query.Where("customer_id = ?", req.CustomerID)
	}

	if err := query.Find(&invoices).Error; err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return invoices, nil
}

func (r *InvoiceRepository) MarkOverdueInvoices(ctx context.Context) error {
	err := r.db.WithContext(ctx).
		Model(&models.Invoice{}).
		Where("due_date < ?", time.Now()).
		Where("paid_amount < amount").
		Where("status IN ?", []string{
			string(constants.Unpaid),
			string(constants.PartiallyPaid),
		}).
		Updates(map[string]any{
			"status":     constants.Overdue,
			"updated_at": time.Now(),
		}).Error

	if err != nil {
		errWrap.WrapError(errConstant.ErrSQLError)
	}

	return nil
}
