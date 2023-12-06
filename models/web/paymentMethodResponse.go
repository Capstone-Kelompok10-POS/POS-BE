package web

type PaymentMethodResponse struct {
	ID            uint                `json:"id"`
	PaymentTypeID uint                `json:"paymentTypeId"`
	PaymentType   PaymentTypeResponse `json:"paymentType"`
	Name          string              `json:"name"`
}
