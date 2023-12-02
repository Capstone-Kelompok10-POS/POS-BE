package repository

import (
	"gorm.io/gorm"
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"
)

type ProductDetailRepository interface {
	Create(productDetail *domain.ProductDetail) (*domain.ProductDetail, error)
	Update(productDetail *domain.ProductDetail, id uint) (*domain.ProductDetail, error)
	FindById(id uint) (*domain.ProductDetail, error)
	FindAll() ([]domain.ProductDetail, error)
	Delete(id uint) error
}

type ProductDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductDetailRepository(DB *gorm.DB) ProductDetailRepository {
	return &ProductDetailRepositoryImpl{DB: DB}
}

func (repository *ProductDetailRepositoryImpl) Create(request *domain.ProductDetail) (*domain.ProductDetail, error) {
	ProductDB := req.ProductDetailDomainToProductDetailSchema(*request)

	result := repository.DB.Create(&ProductDB)

	if result.Error != nil {
		return nil, result.Error
	}

	response := res.ProductDetailSchemaToProductDetailDomain(ProductDB)

	return response, nil
}

func (repository *ProductDetailRepositoryImpl) Update(productDetail *domain.ProductDetail, id uint) (*domain.ProductDetail, error) {
	result := repository.DB.Table("products_detail").Where("id = ?", id).Updates(productDetail)

	if result.Error != nil {
		return nil, result.Error
	}

	return productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) FindById(id uint) (*domain.ProductDetail, error) {
	product := domain.ProductDetail{}

	result := repository.DB.Where("deleted_at IS NULL").First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *ProductDetailRepositoryImpl) FindAll() ([]domain.ProductDetail, error) {
	product := []domain.ProductDetail{}

	result := repository.DB.Where("deleted_at IS NULL").Find(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repository *ProductDetailRepositoryImpl) Delete(id uint) error {
	result := repository.DB.Delete(&schema.ProductDetail{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
