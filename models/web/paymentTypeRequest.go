package web

type PaymentTypeRequest struct {
	TypeName string `json:"typeName" validate:"required"`
}
