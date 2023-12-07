package web

type TransactionPaymentCreateRequest struct {
	PaymentMethodID uint   `json:"paymentMethodId"`
	PaymentTypeID   uint   `json:"paymentTypeId"`
	// Invoice         string `json:"invoice"`
}
