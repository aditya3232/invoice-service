package error

func ErrMapping(err error) bool {
	var (
		GeneralErrors = GeneralErrors
		InvoiceErrors = InvoiceErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, InvoiceErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
