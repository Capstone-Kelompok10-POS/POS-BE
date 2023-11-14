package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func CashierCreateRequestToCashierDomain(request web.CashierCreateRequest) *domain.Cashier {
	return &domain.Cashier{
		AdminID: request.AdminID,
		Fullname: request.Fullname,
		Username: request.Username,
		Password: request.Password,
	}
}

func CashierLoginRequestToCashierDomain(request web.CashierLoginRequest) *domain.Cashier {
	return &domain.Cashier{
		Username: request.Username,
		Password: request.Password,
	}
}

func CashierDomainintoCashierSchema(request domain.Cashier) *schema.Cashier {
	return &schema.Cashier{
		AdminID: request.AdminID,
		Fullname: request.Fullname,
		Username: request.Username,
		Password: request.Password,
	}
}

func CashierUpdateRequestToCashierDomain(request web.CashierUpdateRequest) *domain.Cashier {
	return &domain.Cashier{
		AdminID: request.AdminID,
		Fullname: request.Fullname,
		Username: request.Username,
		Password: request.Password,
	}
}
