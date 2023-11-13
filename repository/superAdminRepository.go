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
	SuperAdmin := domain.SuperAdmin{}

	result := repository.DB.First(&SuperAdmin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SuperAdmin, nil
}

func (repository *SuperAdminRepositoryImpl) FindByUsername(username string) (*domain.SuperAdmin, error) {
	SuperAdmin := domain.SuperAdmin{}

	result := repository.DB.Where("username = ?", username).First(&SuperAdmin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &SuperAdmin, nil
}

func (repository *SuperAdminRepositoryImpl) FindAll() ([]domain.SuperAdmin, error) {
	SuperAdmin := []domain.SuperAdmin{}

	result := repository.DB.Find(&SuperAdmin)
	if result.Error != nil {
		return nil, result.Error
	}
	return SuperAdmin, nil
}

