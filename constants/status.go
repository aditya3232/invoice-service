package constants

type InvoiceStatusString string

const (
	Unpaid        InvoiceStatusString = "UNPAID"
	PartiallyPaid InvoiceStatusString = "PARTIALLY_PAID"
	Paid          InvoiceStatusString = "PAID"
	Overdue       InvoiceStatusString = "OVERDUE"
)

func (s InvoiceStatusString) IsValid() bool {
	switch s {
	case Unpaid, PartiallyPaid, Paid, Overdue:
		return true
	default:
		return false
	}
}
