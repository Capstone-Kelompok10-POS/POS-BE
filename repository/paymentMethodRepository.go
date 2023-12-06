package repository

import (
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	Create(paymentMethod *domain.PaymentMethod) (*domain.PaymentMethod, error)
	Update(paymentMethod *domain.PaymentMethod, id int) (*domain.PaymentMethod, error)
	FindById(id int) (*domain.PaymentMethod, error)
	FindByName(name string) (*domain.PaymentMethod, error)
	FindAll() ([]domain.PaymentMethod, error)
	Delete(id int) error
}

type PaymentMethodRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentMethodRepository(DB *gorm.DB) PaymentMethodRepository {
	return &PaymentMethodRepositoryImpl{DB: DB}
}

func (repository *PaymentMethodRepositoryImpl) Create(paymentMethod *domain.PaymentMethod) (*domain.PaymentMethod, error) {
	paymentMethodDB := req.PaymentMethodDomainToPaymentMethodSchema(paymentMethod)
	result := repository.DB.Create(&paymentMethodDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.PaymentMethodSchemaToPaymentMethodDomain(paymentMethodDB)

	return results, nil
}

func (repository *PaymentMethodRepositoryImpl) Update(paymentMethod *domain.PaymentMethod, id int) (*domain.PaymentMethod, error) {
	result := repository.DB.Table("payment_methods").Where("id = ?", id).Updates(domain.PaymentMethod{
		ID:            paymentMethod.ID,
		PaymentTypeID: paymentMethod.PaymentTypeID,
		Name:          paymentMethod.Name,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentMethod, nil
}

func (repository *PaymentMethodRepositoryImpl) FindById(id int) (*domain.PaymentMethod, error) {
	paymentMethod := domain.PaymentMethod{}

	result := repository.DB.First(&paymentMethod, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &paymentMethod, nil
}

func (repository *PaymentMethodRepositoryImpl) FindByName(name string) (*domain.PaymentMethod, error) {
	PaymentMethod := domain.PaymentMethod{}

	result := repository.DB.Where("name LIKE ?", "%"+name+"%").First(&PaymentMethod)
	if result.Error != nil {
		return nil, result.Error
	}

	return &PaymentMethod, nil
}

func (repository *PaymentMethodRepositoryImpl) FindAll() ([]domain.PaymentMethod, error) {
	paymentMethod := []domain.PaymentMethod{}
	query := "SELECT * FROM payment_methods WHERE deleted_at IS NULL"
	result := repository.DB.Raw(query).Scan(&paymentMethod)
	if result.Error != nil {
		return nil, result.Error
	}
	return paymentMethod, nil
}

func (repository *PaymentMethodRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.PaymentMethod{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
