package repository

import (
	"fmt"
	"qbills/models/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	BeginTransaction() *gorm.DB
	Save(transaction *domain.Transaction) (*domain.Transaction, error)
	FindById(transactionID int) (*domain.Transaction, error)
	// UpdateTransactionStatus(transaction domain.) (*domain.Transaction, error)
}

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{DB: DB}
}

func (repository *TransactionRepositoryImpl) BeginTransaction() *gorm.DB {
    return repository.DB.Begin()
}

func (repository *TransactionRepositoryImpl) Save(transaction *domain.Transaction) (*domain.Transaction, error) {
	tx := repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
    result := tx.Create(&transaction)

    if result.Error != nil {
        tx.Rollback()
        return nil, result.Error
    }

    // Commit if transaction successfully
    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return nil, fmt.Errorf("error committing transaction: %w", err)
    }

    return transaction, nil
}

func (repository *TransactionRepositoryImpl) FindById(transactionID int) (*domain.Transaction, error) {
	transaction := domain.Transaction{}

	result := repository.DB.Preload("Cashier").Preload("Membership").Preload("ConvertPoint").Preload("Details").First(&transaction, transactionID)
	if result.Error != nil {
		return nil, result.Error
	}

	// Preload the items and the category fields
	for i := range transaction.Details {
		if err := repository.DB.Preload("Product").Model(&transaction.Details[i]).Association("ProductDetail").Find(&transaction.Details[i].ProductDetail); err != nil {
			return nil, err
		}
	}

	return &transaction, nil
}



