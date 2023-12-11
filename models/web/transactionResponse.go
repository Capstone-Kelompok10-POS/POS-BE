package web

type TransactionResponse struct {
	ID                 uint                          `json:"id"`
	CreatedAt          string                        `json:"createdAt"`
	Cashier            CashierTransactionResponse    `json:"cashier"`
	Membership         MembershipTransactionResponse `json:"membership"`
	ConvertPointID     uint                          `json:"convertPointId"`
	Discount           float64                       `json:"discount"`
	TotalPrice         float64                       `json:"totalPrice"`
	Tax                float64                       `json:"tax"`
	TotalPayment       float64                       `json:"totalPayment"`
	Details            []TransactionDetailResponse   `json:"details"`
	TransactionPayment TransactionPaymentResponse    `json:"transactionPayment"`
}

type TransactionMonthlyRevenueResponse struct {
	Year    int     `json:"year"`
	Month   int     `json:"month"`
	Revenue float64 `json:"revenue"`
}
type TransactionYearlyRevenueResponse struct {
	Year    int     `json:"year"`
	Revenue float64 `json:"revenue"`
}