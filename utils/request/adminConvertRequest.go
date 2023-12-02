package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func AdminCreateRequestToAdminDomain(request web.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		SuperAdminID: request.SuperAdminID,
		FullName:       request.FullName,
		Username:       request.Username,
		Password:       request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request web.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Username: request.Username,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request web.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		FullName:       request.FullName,
		Username:       request.Username,
		Password:       request.Password,
	}
}

func AdminDomainToAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		SuperAdminID: request.SuperAdminID,
		FullName:       request.FullName,
		Username:       request.Username,
		Password:       request.Password,
	}
}