package domain

type PaymentMethod struct {
	ID            uint
	PaymentTypeID uint
	PaymentType   PaymentType
	Name          string
}
