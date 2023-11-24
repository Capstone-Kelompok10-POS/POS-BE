package repository

import (
	"gorm.io/gorm"
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"
)

type PaymentTypeRepository interface {
	Create(paymentType *domain.PaymentType) (*domain.PaymentType, error)
	Update(paymentType *domain.PaymentType, id int) (*domain.PaymentType, error)
	FindById(id int) (*domain.PaymentType, error)
	FindByName(name string) (*domain.PaymentType, error)
	FindAll() ([]domain.PaymentType, error)
	Delete(id int) error
}

type PaymentTypeRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentTypeRepository(DB *gorm.DB) PaymentTypeRepository {
	return &PaymentTypeRepositoryImpl{DB: DB}
}

func (repository *PaymentTypeRepositoryImpl) Create(paymentType *domain.PaymentType) (*domain.PaymentType, error) {
	paymentTypeDB := req.PaymentTypeDomainToPaymentTypeSchema(*paymentType)
	result := repository.DB.Create(&paymentTypeDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.PaymentTypeSchemaToPaymentTypeDomain(paymentTypeDB)

	return results, nil
}

func (repository *PaymentTypeRepositoryImpl) Update(paymentType *domain.PaymentType, id int) (*domain.PaymentType, error) {
	result := repository.DB.Table("payment_types").Where("id = ?", id).Updates(domain.PaymentType{
		ID:       paymentType.ID,
		TypeName: paymentType.TypeName,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentType, nil
}

func (repository *PaymentTypeRepositoryImpl) FindById(id int) (*domain.PaymentType, error) {
	paymentType := domain.PaymentType{}

	result := repository.DB.First(&paymentType, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &paymentType, nil
}

func (repository *PaymentTypeRepositoryImpl) FindByName(name string) (*domain.PaymentType, error) {
	paymentType := domain.PaymentType{}

	result := repository.DB.Where("name = ?", name).First(&paymentType)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentType, nil
}

func (repository *PaymentTypeRepositoryImpl) FindAll() ([]domain.PaymentType, error) {
	paymentType := []domain.PaymentType{}

	result := repository.DB.Find(&paymentType)
	if result.Error != nil {
		return nil, result.Error
	}
	return paymentType, nil
}

func (repository *PaymentTypeRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.PaymentType{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
