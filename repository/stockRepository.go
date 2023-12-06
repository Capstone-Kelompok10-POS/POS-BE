package repository

import (
	"gorm.io/gorm"
	"qbills/models/domain"
	"qbills/utils/request"
	"qbills/utils/response"
)

type StockRepository interface {
	Create(stock *domain.Stock) (*domain.Stock, error)
	FindAll() ([]domain.Stock, error)
	FindById(id uint) (*domain.Stock, error)
	FindIncreaseStock() ([]domain.Stock, error)
	FindDecreaseStock() ([]domain.Stock, error)
}

type StockRepositoryImpl struct {
	DB *gorm.DB
}

func NewStockRepository(DB *gorm.DB) StockRepository {
	return &StockRepositoryImpl{DB: DB}
}

func (repository *StockRepositoryImpl) Create(stock *domain.Stock) (*domain.Stock, error) {

	req := request.StockDomainToStockSchema(*stock)

	result := repository.DB.Create(&req)

	if result.Error != nil {
		return nil, result.Error
	}

	res := response.StockSchemaToStockDomain(req)

	return res, nil
}

func (repository *StockRepositoryImpl) FindAll() ([]domain.Stock, error) {
	stock := []domain.Stock{}

	if err := repository.DB.Preload("ProductDetail").Find(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

func (repository *StockRepositoryImpl) FindById(id uint) (*domain.Stock, error) {
	stock := domain.Stock{}

	result := repository.DB.Preload("ProductDetail").Where("deleted_at IS NULL").First(&stock, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &stock, nil
}

func (repository *StockRepositoryImpl) FindIncreaseStock() ([]domain.Stock, error) {
	stock := []domain.Stock{}

	if err := repository.DB.Where("stock > 0").Preload("ProductDetail").Find(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

func (repository *StockRepositoryImpl) FindDecreaseStock() ([]domain.Stock, error) {
	stock := []domain.Stock{}

	if err := repository.DB.Where("stock < 0").Preload("ProductDetail").Find(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}
