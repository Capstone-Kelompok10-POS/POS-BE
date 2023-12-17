package domain

import "time"

type Transaction struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt time.Time
	CashierID      uint
	Cashier        Cashier
	MembershipID   uint
	Membership     Membership
	ConvertPointID uint
	ConvertPoint   ConvertPoint

	Discount     float64
	TotalPrice   float64
	Tax          float64
	TotalPayment float64
	Details      []TransactionDetail
	TransactionPayment TransactionPayment
}

type TransactionMonthlyRevenue struct{
	Year         int     
	Month        int     
	Revenue 	 float64 
}
type TransactionYearlyRevenue struct{
	Year    int    
	Revenue float64 
}

type TransactionDailyRevenue struct{
	Day time.Time
	Success int
	Pending int
	Cancelled int 
	Revenue float64 
}