package web

type PaymentMethodRequest struct {
	PaymentTypeID uint   `json:"paymentTypeId" validate:"required,number"`
	Name          string `json:"name" validate:"required"`
}
