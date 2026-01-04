package error

import "errors"

var (
	ErrInvoiceNotFound            = errors.New("invoice not found")
	ErrCustomerNotFound           = errors.New("customer not found")
	ErrMarkOverdueAlreadyFullPaid = errors.New("invoice already fully paid")
	ErrMarkOverdueNotOverdueYet   = errors.New("invoice is not overdue yet")
	ErrMarkOverdueInvalidStatus   = errors.New("invoice status is not eligible to be marked as overdue")
)

var InvoiceErrors = []error{
	ErrInvoiceNotFound,
	ErrCustomerNotFound,
	ErrMarkOverdueAlreadyFullPaid,
	ErrMarkOverdueNotOverdueYet,
	ErrMarkOverdueInvalidStatus,
}
