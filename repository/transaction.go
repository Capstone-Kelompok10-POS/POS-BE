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
	FindRecentTransaction() ([]domain.Transaction, error)
	FindMonthlyRevenue(currentYear int) ([]domain.TransactionMonthlyRevenue, error)
	FindYearlyRevenue(currentYear int) (*domain.TransactionYearlyRevenue, error)
	FindDailyTransaction(currentDate string) (*domain.TransactionDailyRevenue, error)
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

func (repository *TransactionRepositoryImpl) GetTotalPayment() (float64, error) {
	var totalPayment float64

	result := repository.DB.Table("transactions").
		Joins("JOIN transaction_payments ON transactions.id = transaction_payments.transaction_id").
		Where("transactions.deleted_at IS NULL").
		Where("transaction_payments.payment_status = 'success'").
		Select("SUM(transaction_payments.total_payment)").
		Scan(&totalPayment)

	if result.Error != nil {
		return 0, result.Error
	}

	return totalPayment, nil
}

func (repository *TransactionRepositoryImpl) FindRecentTransaction() ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}

	result := repository.DB.Preload("Cashier").Preload("Membership").Preload("ConvertPoint").Preload("Details.ProductDetail.Product").Preload("Details.ProductDetail").Preload("TransactionPayment.PaymentMethod").Preload("TransactionPayment.PaymentMethod.PaymentType").Where("transactions.deleted_at IS NULL ORDER BY created_at DESC LIMIT 6").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions,  nil
}

func (repository *TransactionRepositoryImpl) FindMonthlyRevenue(currentYear int) ([]domain.TransactionMonthlyRevenue, error) {
	monthlyTotalPayments := []domain.TransactionMonthlyRevenue{}

	query := `
	SELECT YEAR(t.created_at) as year, MONTH(t.created_at) as month, SUM(t.total_payment) as revenue
	FROM transactions t
	JOIN transaction_payments tp ON t.id = tp.transaction_id
	WHERE t.deleted_at IS NULL AND tp.payment_status = 'success' AND YEAR(t.created_at) = ?
	GROUP BY YEAR(t.created_at), MONTH(t.created_at)
	ORDER BY year, month
	`

	result := repository.DB.Raw(query, currentYear).Scan(&monthlyTotalPayments)

	if result.Error != nil {
		return nil, result.Error
	}

	return monthlyTotalPayments, nil
}

func (repository *TransactionRepositoryImpl) FindYearlyRevenue(currentYear int) (*domain.TransactionYearlyRevenue, error) {
	yearlyRevenue := domain.TransactionYearlyRevenue{}

	query := `
	SELECT YEAR(t.created_at) as year, SUM(t.total_payment) as revenue
		FROM transactions t
		JOIN transaction_payments tp ON t.id = tp.transaction_id
		WHERE t.deleted_at IS NULL AND tp.payment_status = 'success' AND YEAR(t.created_at) = ?
		GROUP BY YEAR(t.created_at)
		ORDER BY year
	`

	result := repository.DB.Raw(query, currentYear).Scan(&yearlyRevenue)

	if result.Error != nil {
		return nil, result.Error
	}

	return &yearlyRevenue, nil
}

func (repository *TransactionRepositoryImpl) FindDailyTransaction(currentDate string) (*domain.TransactionDailyRevenue, error) {
	dailyRevenue := domain.TransactionDailyRevenue{}

	query := `
	SELECT
		DATE(t.created_at) as day,
		COUNT(CASE WHEN tp.payment_status = 'success' THEN 1 END) as success,
		COUNT(CASE WHEN tp.payment_status = 'pending' THEN 1 END) as pending,
		COUNT(CASE WHEN tp.payment_status = 'failure' THEN 1 END) as cancelled,
		SUM(CASE WHEN tp.payment_status = 'success' THEN t.total_payment ELSE 0 END) as revenue
	FROM transactions t
	LEFT JOIN transaction_payments tp ON t.id = tp.transaction_id
	WHERE t.deleted_at IS NULL AND DATE(t.created_at) = ?
	GROUP BY day
	`

	result := repository.DB.Raw(query, currentDate).Scan(&dailyRevenue)

	if result.Error != nil {
		return nil, result.Error
	}
	
	if dailyRevenue.Success == 0 && dailyRevenue.Pending == 0 && dailyRevenue.Cancelled == 0 && dailyRevenue.Revenue == 0 {
		return nil, fmt.Errorf("transaction daily not found")
	}

	return &dailyRevenue, nil
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