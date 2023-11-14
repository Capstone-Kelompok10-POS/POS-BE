package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func CashierDomainToCashierLoginResponse(cashier *domain.Cashier) web.CashierLoginResponse {
	return web.CashierLoginResponse{
		Fullname: cashier.Fullname,
		Username: cashier.Username,
	}
}

func CashierSchemaToCashierDomain(cashier *schema.Cashier) *domain.Cashier {
	return &domain.Cashier{
		ID:       cashier.ID,
		AdminID: cashier.AdminID,
		Fullname: cashier.Fullname,
		Username: cashier.Username,
	}
}

func CashierDomainToCashierResponse(cashier *domain.Cashier) web.CashierResponse {
	return web.CashierResponse{
		ID:       cashier.ID,
		AdminID: cashier.AdminID,
		Fullname: cashier.Fullname,
		Username: cashier.Username,
	}
}

func ConvertCashierResponse(cashiers []domain.Cashier) []web.CashierResponse {
	var results []web.CashierResponse
	for _, cashier := range cashiers {
		cashierResponse := web.CashierResponse{
			ID: cashier.ID,
			AdminID: cashier.AdminID,
			Fullname: cashier.Fullname,
			Username: cashier.Username,
		}
		results = append(results, cashierResponse)
	}
	return results
}

