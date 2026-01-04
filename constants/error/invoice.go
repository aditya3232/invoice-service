package error

import "errors"

var (
	ErrInvoiceNotFound  = errors.New("invoice not found")
	ErrCustomerNotFound = errors.New("customer not found")
)

var InvoiceErrors = []error{
	ErrInvoiceNotFound,
	ErrCustomerNotFound,
}
