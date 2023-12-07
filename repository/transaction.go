package repository

import (
	"fmt"
	"os"
	"qbills/models/domain"
	"qbills/utils/helpers"
	"strconv"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	BeginTransaction() *gorm.DB
	Save(transaction *domain.Transaction) (*domain.Transaction, error)
	FindById(transactionID int) (*domain.Transaction, error)
	UpdateStatusTransactionPayment(status string, transactionResult *domain.PaymentTransactionStatus) error
	FindAllTransaction() ([]domain.Transaction, int, error)
	FindByInvoice(invoice string) (*domain.Transaction, error)
	FindPaginationTransaction(orderBy string, paginate helpers.Pagination) ([]domain.Transaction, *helpers.Pagination, error)
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

	result := repository.DB.Preload("Cashier").Preload("Membership").Preload("ConvertPoint").Preload("Details").Preload("TransactionPayment.PaymentMethod").Preload("TransactionPayment.PaymentMethod.PaymentType").First(&transaction, transactionID)
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

func (repository *TransactionRepositoryImpl) UpdateStatusTransactionPayment(status string, transactionResult *domain.PaymentTransactionStatus) error {
	result := repository.DB.Table("transaction_payments").Where("invoice = ?", transactionResult.OrderID).Updates(&domain.TransactionPayment{
		PaymentStatus: status,
		UpdatedAt: transactionResult.SettlementTime,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *TransactionRepositoryImpl) FindByInvoice(invoice string) (*domain.Transaction, error) {
	transaction := domain.Transaction{}

	result := repository.DB.
		Preload("Cashier").
		Preload("Membership").
		Preload("ConvertPoint").
		Preload("Details.ProductDetail.Product").
		Preload("Details.ProductDetail").
		Preload("TransactionPayment.PaymentMethod").
		Preload("TransactionPayment.PaymentMethod.PaymentType").
		Joins("LEFT JOIN transaction_payments ON transactions.id = transaction_payments.transaction_id").
		Where("transactions.deleted_at IS NULL AND transaction_payments.invoice = ?", invoice).Find(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}


	return &transaction, nil
}

func (repository *TransactionRepositoryImpl) FindAllTransaction() ([]domain.Transaction, int, error) {
	transactions := []domain.Transaction{}

	result := repository.DB.Preload("Cashier").Preload("Membership").Preload("ConvertPoint").Preload("Details.ProductDetail.Product").Preload("Details.ProductDetail").Preload("TransactionPayment.PaymentMethod").Preload("TransactionPayment.PaymentMethod.PaymentType").Where("deleted_at IS NULL").Find(&transactions)
	if result.Error != nil {
		return nil, 0,  result.Error
	}

	totalTransaction := len(transactions)

	return transactions, totalTransaction , nil
}

func (repository *TransactionRepositoryImpl) FindPaginationTransaction(orderBy string, paginate helpers.Pagination) ([]domain.Transaction, *helpers.Pagination, error) {
	var transactions []domain.Transaction

	result := repository.DB.Scopes(helpers.Paginate(transactions, &paginate, repository.DB)).Preload("Cashier").Preload("Membership").Preload("ConvertPoint").Preload("Details.ProductDetail.Product").Preload("Details.ProductDetail").Preload("TransactionPayment.PaymentMethod").Preload("TransactionPayment.PaymentMethod.PaymentType").Where("deleted_at IS NULL")

	if orderBy != "" {
		result.Order("name" + orderBy).Find(&transactions)
	} else {
		result.Find(&transactions)
	}

	if result.Error != nil {
		return nil, nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil, fmt.Errorf("transactions is empty")
	}

	if paginate.Page <= 1 {
		paginate.PreviousPage = ""
	} else {
		paginate.PreviousPage = os.Getenv("MAIN_URL") + "/api/" + os.Getenv("API_VERSION") + "/transactions/pagination?page=" + strconv.Itoa(int(paginate.Page)-1)
	}

	if paginate.Page >= paginate.TotalPage {
		paginate.NextPage = ""
	} else {
		paginate.NextPage = os.Getenv("MAIN_URL") + "/api/" + os.Getenv("API_VERSION") + "/transactions/pagination?page=" + strconv.Itoa(int(paginate.Page)+1)
	}

	return transactions, &paginate, nil
}