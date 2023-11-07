package repository

import (
	"gorm.io/gorm"
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"
)

type ProductTypeRepository interface {
	Create(productType *domain.ProductType) (*domain.ProductType, error)
	Update(productType *domain.ProductType, id uint) (*domain.ProductType, error)
	FindById(id uint) (*domain.ProductType, error)
	FindAll() ([]domain.ProductType, error)
	FindByName(name string) (*domain.ProductType, error)
	Delete(id uint) error
}

type ProductTypeRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductTypeRepository(DB *gorm.DB) ProductTypeRepository {
	return &ProductTypeRepositoryImpl{DB: DB}
}

func (repository *ProductTypeRepositoryImpl) Create(productType *domain.ProductType) (*domain.ProductType, error) {
	productTypeDB := req.ProductTypeDomainToProductTypeSchema(*productType)

	result := repository.DB.Create(&productTypeDB)

	if result.Error != nil {
		return nil, result.Error
	}

	results := res.ProductTypeSchemaToProductTypeDomain(productTypeDB)

	return results, nil
}

func (repository *ProductTypeRepositoryImpl) Update(productType *domain.ProductType, id uint) (*domain.ProductType, error) {
	result := repository.DB.Table("ProductType").Where("id = ?", id).Updates(domain.ProductType{
		TypeName:        productType.TypeName,
		TypeDescription: productType.TypeDescription,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return productType, nil
}

func (repository *ProductTypeRepositoryImpl) FindById(id uint) (*domain.ProductType, error) {
	productType := domain.ProductType{}

	result := repository.DB.First(&productType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &productType, nil
}

func (repository *ProductTypeRepositoryImpl) FindAll() ([]domain.ProductType, error) {
	productType := []domain.ProductType{}

	result := repository.DB.Find(&productType)

	if result.Error != nil {
		return nil, result.Error

	}

	return productType, nil
}

func (repository *ProductTypeRepositoryImpl) FindByName(name string) (*domain.ProductType, error) {
	productType := domain.ProductType{}

	result := repository.DB.Where("LOWER(typeName) LIKE LOWER(?)", "%"+name+"%").First(&productType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &productType, nil
}

func (repository *ProductTypeRepositoryImpl) Delete(id uint) error {
	result := repository.DB.Delete(&schema.ProductType{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
