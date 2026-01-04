package dto

type PaymentData struct {
	PaymentID   int     `json:"payment_id"`
	InvoiceID   int     `json:"invoice_id"`
	Amount      float64 `json:"amount"`
	ReferenceNo string  `json:"reference_no"`
}

type PaymentContent struct {
	Event    KafkaEvent             `json:"event"`
	Metadata KafkaMetaData          `json:"metadata"`
	Body     KafkaBody[PaymentData] `json:"body"`
}
