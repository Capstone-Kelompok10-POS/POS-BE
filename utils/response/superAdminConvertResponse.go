package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func SuperAdminDomainToSuperAdminLoginResponse(superadmin *domain.SuperAdmin) web.SuperAdminLoginResponse {
	return web.SuperAdminLoginResponse{
		Username: superadmin.Username,
	}
}

func SuperAdminSchemaToSuperAdminDomain(superAdminadmin *schema.SuperAdmin) *domain.SuperAdmin {
	return &domain.SuperAdmin{
		ID:             superAdminadmin.ID,
		Username:       superAdminadmin.Username,
	}
}

func SuperAdminDomainToSuperAdminResponse(superAdmin *domain.SuperAdmin) web.SuperAdminResponse {
	return web.SuperAdminResponse{
		ID:             superAdmin.ID,
		Username:       superAdmin.Username,
	}
}

func ConvertSuperAdminResponse(superAdmins []domain.SuperAdmin) []web.SuperAdminResponse {
	var results []web.SuperAdminResponse
	for _, superAdmin := range superAdmins {
		superAdminResponse := web.SuperAdminResponse{
			ID:             superAdmin.ID,
			Username:       superAdmin.Username,
		}
		results = append(results, superAdminResponse)
	}
	return results
}