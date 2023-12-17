package repository

import (
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type ProductDetailRepository interface {
	Create(productDetail *domain.ProductDetail) (*domain.ProductDetail, error)
	Update(productDetail *domain.ProductDetail, id uint) (*domain.ProductDetail, error)
	Save(productDetail *domain.ProductDetail, id uint) (*domain.ProductDetail, error)
	StockDecrease(tx *gorm.DB, productDetail *domain.ProductDetail) error
	FindById(id uint) (*domain.ProductDetail, error)
	FindByProductId(ProductID uint) ([]domain.ProductDetail, error)
	FindAll() ([]domain.ProductDetail, error)
	Delete(id uint) error
	FindAllByIds(ids []uint) ([]domain.ProductDetail, error)
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
	result := repository.DB.Table("product_details").Where("id = ?", id).Updates(productDetail)

	if result.Error != nil {
		return nil, result.Error
	}

	return productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) Save(productDetail *domain.ProductDetail, id uint) (*domain.ProductDetail, error) {
	result := repository.DB.Table("product_details").Where("id = ?", id).Save(productDetail)

	if result.Error != nil {
		return nil, result.Error
	}

	return productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) StockDecrease(tx *gorm.DB, productDetail *domain.ProductDetail) error {
	result := tx.Table("product_details").Where("id = ?", productDetail.ID).Where("deleted_at IS NULL").Update("total_stock", productDetail.TotalStock)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *ProductDetailRepositoryImpl) FindById(id uint) (*domain.ProductDetail, error) {
	productDetail := domain.ProductDetail{}
	result := repository.DB.Where("deleted_at IS NULL").First(&productDetail, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &productDetail, nil
}

func (repository *ProductDetailRepositoryImpl) FindByProductId(ProductID uint) ([]domain.ProductDetail, error) {
	productsDetail := []domain.ProductDetail{}

	result := repository.DB.Where("product_id = ?", ProductID).Find(&productsDetail)

	if result.Error != nil {
		return nil, result.Error
	}

	return productsDetail, nil
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
	result := repository.DB.Where("deleted_at IS NULL").Delete(&schema.ProductDetail{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *ProductDetailRepositoryImpl) FindAllByIds(ids []uint) ([]domain.ProductDetail, error) {
	var products []domain.ProductDetail

	result := repository.DB.Preload("Product").Where("deleted_at IS NULL").Find(&products, ids)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
