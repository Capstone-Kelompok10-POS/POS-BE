package repository

import (
	"qbills/models/domain"

	"gorm.io/gorm"
)

type SuperAdminRepository interface {
	FindById(id int) (*domain.SuperAdmin, error)
	FindByUsername(username string) (*domain.SuperAdmin, error)
	FindAll() ([]domain.SuperAdmin, error)
}

type SuperAdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewSuperAdminRepository(DB *gorm.DB) SuperAdminRepository {
	return &SuperAdminRepositoryImpl{DB: DB}
}


func (repository *SuperAdminRepositoryImpl) FindById(id int) (*domain.SuperAdmin, error) {
	superAdmin := domain.SuperAdmin{}

	result := repository.DB.First(&superAdmin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &superAdmin, nil
}

func (repository *SuperAdminRepositoryImpl) FindByUsername(username string) (*domain.SuperAdmin, error) {
	superAdmin := domain.SuperAdmin{}

	query := "SELECT super_admins.* FROM super_admins WHERE LOWER(username) = LOWER(?) AND deleted_at IS NULL"

	result := repository.DB.Raw(query, username).Scan(&superAdmin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &superAdmin, nil
}

func (repository *SuperAdminRepositoryImpl) FindAll() ([]domain.SuperAdmin, error) {
	superAdmin := []domain.SuperAdmin{}

	result := repository.DB.Find(&superAdmin)
	if result.Error != nil {
		return nil, result.Error
	}
	return superAdmin, nil
}

