package web

type TransactionPaymentCreateRequest struct {
	PaymentMethodID uint `json:"paymentMethodId"`
	PaymentTypeID   uint `json:"paymentTypeId"`
}

type TransactionPaymentUpdateRequest struct {
	Invoice string `json:"invoice"`
}
