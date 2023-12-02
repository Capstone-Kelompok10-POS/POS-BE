package web

type PaymentMethodResponse struct {
	ID            uint   `json:"id"`
	PaymentTypeID uint   `json:"paymentTypeId"`
	Name          string `json:"name"`
}
