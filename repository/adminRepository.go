package repository

import (
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	Update(admin *domain.Admin, id int) (*domain.Admin, error)
	FindById(id int) (*domain.Admin, error)
	FindByUsername(username string) (*domain.Admin, error)
	FindAll() ([]domain.Admin, error)
	Delete(id int) error
}

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{DB: DB}
}

func (repository *AdminRepositoryImpl) Create(admin *domain.Admin) (*domain.Admin, error) {
	adminDB := req.AdminDomainToAdminSchema(*admin)
	result := repository.DB.Create(&adminDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.AdminSchemaToAdminDomain(adminDB)

	return results, nil
}

func (repository *AdminRepositoryImpl) Update(admin *domain.Admin, id int) (*domain.Admin, error) {
	result := repository.DB.Table("admins").Where("id = ?", id).Updates(domain.Admin{
		FullName: admin.FullName,
		Username: admin.Username,
		Password: admin.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) FindById(id int) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.First(&admin, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindByUsername(username string) (*domain.Admin, error) {
	admin := domain.Admin{}

	query := "SELECT admins.* FROM admins WHERE LOWER(username) = LOWER(?) AND deleted_at IS NULL"
	result := repository.DB.Raw(query, username).Scan(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}


func (repository *AdminRepositoryImpl) FindAll() ([]domain.Admin, error) {
	admin := []domain.Admin{}

	result := repository.DB.Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return admin, nil
}


func (repository *AdminRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
