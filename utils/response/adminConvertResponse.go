package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func AdminDomainToAdminLoginResponse(admin *domain.Admin) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Username: admin.Username,
	}
}

func AdminSchemaToAdminDomain(admin *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:           admin.ID,
		SuperAdminID: admin.SuperAdminID,
		FullName:     admin.FullName,
		Username:     admin.Username,
	}
}

func AdminDomainToAdminResponse(admin *domain.Admin) web.AdminResponse {
	return web.AdminResponse{
		ID:           admin.ID,
		SuperAdminID: admin.SuperAdminID,
		FullName:     admin.FullName,
		Username:     admin.Username,
	}
}

func ConvertAdminResponse(admins []domain.Admin) []web.AdminResponse {
	var results []web.AdminResponse
	for _, admin := range admins {
		adminResponse := web.AdminResponse{
			ID:           admin.ID,
			SuperAdminID: admin.SuperAdminID,
			FullName:     admin.FullName,
			Username:     admin.Username,
		}
		results = append(results, adminResponse)
	}
	return results
}
