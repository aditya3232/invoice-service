package error

import "errors"

var (
	ErrInvoiceNotFound = errors.New("invoice not found")
)

var InvoiceErrors = []error{
	ErrInvoiceNotFound,
}
