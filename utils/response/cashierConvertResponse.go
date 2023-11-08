package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func CashierDomainToCashierLoginResponse(cashier *domain.Cashier) web.CashierLoginResponse {
	return web.CashierLoginResponse{
		Username: cashier.Username,
	}
}

func CashierSchemaToCashierDomain(cashier *schema.Cashier) *domain.Cashier {
	return &domain.Cashier{
		ID:       cashier.ID,
		Admin_ID: cashier.Admin_ID,
		Fullname: cashier.Fullname,
		Username: cashier.Username,
	}
}

func CashierDomainToCashierResponse(cashier *domain.Cashier) web.CashierResponse {
	return web.CashierResponse{
		ID:       cashier.ID,
		Admin_ID: cashier.Admin_ID,
		Fullname: cashier.Fullname,
		Username: cashier.Username,
	}
}
