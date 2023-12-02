package web

import "qbills/models/domain"


type TransactionCreateRequest struct {
	CashierID 	   uint 					   `json:"cashierId" validate:"numeric"`
	MembershipID   uint                        `json:"membershipId" validate:"numeric"`
	ConvertPointID uint                        `json:"convertPointId" validate:"numeric"`
	Details        []domain.TransactionDetail  `json:"details" validate:"required"`
}

type TransactionCreate struct {
	CashierID 	 uint 		`json:"cashierId"`
	Status 		 string 	`json:"status"`
	Discount 	 float64 	`json:"discount"`
	TotalPrice   float64    `json:"total"`
	Tax 		 float64 	`json:"tax"`
	TotalPayment float64 	`json:"totalPayment"`
}