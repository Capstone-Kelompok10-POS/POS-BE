package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)



func SuperAdminLoginRequestToSuperAdminDomain(request web.SuperAdminLoginRequest) *domain.SuperAdmin {
	return &domain.SuperAdmin{
		Username: request.Username,
		Password: request.Password,
	}
}



func SuperAdminDomainToSuperAdminSchema(request domain.SuperAdmin) *schema.SuperAdmin {
	return &schema.SuperAdmin{
		Username: request.Username,
		Password: request.Password,
	}
}