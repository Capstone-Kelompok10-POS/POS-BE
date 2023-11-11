package repository

import (
	"gorm.io/gorm"
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"
)

type ProductRepository interface {
	Create(request *domain.Product) (*domain.Product, error)
	Update(request *domain.Product, id uint) (*domain.Product, error)
	FindById(id uint) (*domain.Product, error)
	FindAll() ([]domain.Product, error)
	FindByName(name string) ([]*domain.Product, error)
	Delete(id uint) error
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}

func (repository *ProductRepositoryImpl) Create(request *domain.Product) (*domain.Product, error) {
	ProductDB := req.ProductDomainToProductSchema(*request)

	result := repository.DB.Create(&ProductDB)

	if result.Error != nil {
		return nil, result.Error
	}

	response := res.ProductSchemaToProductDomain(ProductDB)

	return response, nil
}

func (repository *ProductRepositoryImpl) Update(request *domain.Product, id uint) (*domain.Product, error) {
	result := repository.DB.Table("products").Where("id = ?", id).Updates(domain.Product{
		ProductTypeID: request.ProductTypeID,
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		Stock:         request.Stock,
		Size:          request.Size,
		Image:         request.Image,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (repository *ProductRepositoryImpl) FindById(id uint) (*domain.Product, error) {
	product := domain.Product{}

	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL").First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *ProductRepositoryImpl) FindAll() ([]domain.Product, error) {
	product := []domain.Product{}

	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL").Find(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) FindByName(name string) ([]*domain.Product, error) {
	products := []*domain.Product{}

	result := repository.DB.Where("deleted_at IS NULL AND nama_produk LIKE ?", "%"+name+"%").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repository *ProductRepositoryImpl) Delete(id uint) error {
	result := repository.DB.Delete(&schema.Product{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
