package services

import (
	"fmt"
	"math/rand"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	"qbills/utils/helpers/midtrans"
	req "qbills/utils/request"
	res "qbills/utils/response"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransaction(request web.TransactionCreateRequest) (*web.TransactionResponse, error)
	FindById(id int) (*domain.Transaction, error)
	FindByInvoice(invoice string) (*domain.Transaction, error)
	FindByYearly() (*domain.TransactionYearlyRevenue, error)
	FindByMonthly() ([]domain.TransactionMonthlyRevenue, error)
	FindByDaily() (*domain.TransactionDailyRevenue, error)
	FindAllTransaction() ([]domain.Transaction, int, error)
	FindRecentTransaction() ([]domain.Transaction, error)
	FindPaginationTransaction(orderBy, QueryLimit, QueryPage string) ([]domain.Transaction, *helpers.Pagination, error)
	GetPricesAndSubTotal(details []domain.TransactionDetail) (map[uint]float64, map[uint]float64, error)
	ProductStockDecrese(tx *gorm.DB, details []domain.TransactionDetail) error
	CalculateTotalPrice(details []domain.TransactionDetail) (float64, error)
	CalculateDiscount(id int, totalPrice float64) (float64, error)
	SubtractionPoint(tx *gorm.DB , pointId , membershipId uint) (*domain.Membership, error)
	UpdateMemberPoint(tx *gorm.DB, totalPayment float64, membershipID uint) (error)
	CreateInvoice(paymentMethod, paymentType uint) (string, error)
	NotificationPayment(notificationPayload map[string]interface{}) error
	ManualPayment(invoice string) (*domain.Transaction, error)
}

type TransactionImpl struct {
	TransactionRepository repository.TransactionRepository
	ProductDetailRepository repository.ProductDetailRepository
	ConvertPointRepository repository.ConvertPointRepository
	MembershipRepository repository.MembershipRepository
	PaymentMethodRepository repository.PaymentMethodRepository
	MidtransCoreApi midtrans.MidtransCoreApi
	validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductDetailRepository, convertPointRepository repository.ConvertPointRepository, membershipRepository repository.MembershipRepository, paymentMethodRepository repository.PaymentMethodRepository, midtransCoreApi midtrans.MidtransCoreApi, validate *validator.Validate) *TransactionImpl {
	return &TransactionImpl{
		TransactionRepository: transactionRepository,
		ProductDetailRepository: productRepository,
		ConvertPointRepository: convertPointRepository,
		MembershipRepository: membershipRepository,
		PaymentMethodRepository: paymentMethodRepository,
		MidtransCoreApi: midtransCoreApi,
		validate:              validate,
	}
}

func (service *TransactionImpl) CreateTransaction(request web.TransactionCreateRequest) (*web.TransactionResponse, error) {
	tx := service.TransactionRepository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()


	membership, err := service.MembershipRepository.FindById(int(request.MembershipID))
	if err != nil {
		return nil, fmt.Errorf("failed to find membership: %w", err)
	}

	paymentMethod, err := service.PaymentMethodRepository.FindById(int(request.TransactionPayment.PaymentMethodID))
	if err != nil {
		return nil, fmt.Errorf("failed to find payment method: %w", err)
	}
	
	// input product price to transaction detaul and calculate subTotal to transaction detail
	productPrice, subTotal, err := service.GetPricesAndSubTotal(request.Details)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate subTotal: %w", err)
	}
	
	// Calculate total price
	totalPrice, err := service.CalculateTotalPrice(request.Details)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total price: %w", err)
	}
	
	//calculate discount
	discount, err := service.CalculateDiscount(int(request.ConvertPointID), totalPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate discount: %w", err)
	}

	// calculate the tax transaction
	tax := (0.10 * (totalPrice - discount))

	//calculate total payment 
	totalPayment := (totalPrice - discount) + tax

	err = service.MatchingTotalPrice(request.TotalPrice, totalPrice)
	if err != nil {
		return nil, fmt.Errorf("total price does not match: %w", err)
	}

	err = service.MatchingDiscount(request.Discount,discount)
	if err != nil {
		return nil, fmt.Errorf("discount does not match: %w", err)
	}

	err = service.MatchingTax(request.Tax,tax)
	if err != nil {
		return nil, fmt.Errorf("tax does not match: %w", err)
	}

	err = service.MatchingTotalPayment(request.TotalPayment, totalPayment)
	if err != nil {
		return nil, fmt.Errorf("total Payment does not match: %w", err)
	}

	//Create Invoice transaction Payment
	invoice, err := service.CreateInvoice(request.TransactionPayment.PaymentMethodID, request.TransactionPayment.PaymentMethod.PaymentTypeID)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}
	request.TransactionPayment.Invoice = invoice

	//Add status Transaction
	status := "pending"
	//Add status transaction Payment
	request.TransactionPayment.PaymentStatus = status
	request.TransactionPayment.VANumber = paymentMethod.Name
	transaction := req.TransactionCreateRequestToTransactionDomain(request,
		web.TransactionCreate{
			Discount: discount,
			TotalPrice: totalPrice,
			Tax: tax,
			TotalPayment: totalPayment,
		}, web.TransactionDetailCreate{
			ProductPrice : productPrice,
			SubTotal: subTotal,
		})
	var result *domain.Transaction
	if paymentMethod.PaymentTypeID == 3 {
		//MIDTRANS CORE
		chargeRequest := req.CreateTransactionPaymentRequestToMidtransChargeRequest(transaction, membership, paymentMethod, discount, tax)
		chargeRequestResponse, err := service.MidtransCoreApi.ChargeTransaction(chargeRequest)
		if chargeRequestResponse == nil {
			return nil, fmt.Errorf("MidtransCoreApi is not initialized %w", err)
		}
		if err != nil {
			return nil, fmt.Errorf("error when creating transaction payment %w", err)
		}
		
		transactionPayment := req.ChargeResponseToTransactionPayment(chargeRequestResponse, transaction)
		result, err = service.TransactionRepository.Save(transactionPayment)
		if err != nil {
			return nil, fmt.Errorf("error when creating transaction %w", err)
		}

	} else {
		result, err = service.TransactionRepository.Save(transaction)
		if err != nil {
			return nil, fmt.Errorf("error when creating transaction %w", err)
		}
	}

	if result != nil {
		_ , err = service.SubtractionPoint(tx, result.ConvertPointID, result.MembershipID)
		if err != nil {
			return nil, fmt.Errorf("error when decreasing point membership %w", err)
		}

		err = service.UpdateMemberPoint(tx, result.TotalPayment, result.MembershipID)
		if err != nil {
			return nil, fmt.Errorf("error when increasing point membership %w", err)
		}
		
		// decrease product total stock
		err = service.ProductStockDecrese(tx, request.Details)
		if err != nil {
			return nil, fmt.Errorf("failed to decrease product stock: %w", err)
		}

		err = tx.Commit().Error
		if err != nil {
			return nil, fmt.Errorf("error committing transaction: %w", err)
		}

		result, err = service.TransactionRepository.FindById(int(result.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to find transaction: %w", err)
		}
		if result.Membership.ID == 0 {
			response := res.TransactionDomainToTransactionResponseNoMembership(result)
			return response, nil
		} else {
			response := res.TransactionDomainToTransactionResponse(result)
			return response, nil
		}
	} else {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
}

func (service *TransactionImpl) ProductStockDecrese(tx *gorm.DB, details []domain.TransactionDetail) error {
	for _, detail := range details {
		// Fetch the item from the database to get the latest stock
		productDetail, err := service.ProductDetailRepository.FindById(detail.ProductDetailID)
		if err != nil {
			return fmt.Errorf("failed to find product: %w", err)
		}
		if productDetail.TotalStock < detail.Quantity {
			return fmt.Errorf("insufficient stock for product with ID: %d", productDetail.ID)
		}
		if productDetail.TotalStock <= 0 {
			return fmt.Errorf("insufficient stock for product with ID: %d", productDetail.ID)
		}

		productDetail.TotalStock -= detail.Quantity

		err = service.ProductDetailRepository.StockDecrease(tx, productDetail)
		if err != nil {
			return fmt.Errorf("failed to update product stock: %w", err)
		}
	}

	return nil
}

func (service *TransactionImpl) GetPricesAndSubTotal(details []domain.TransactionDetail) (map[uint]float64, map[uint]float64, error) {
	productIDs := make([]uint, 0)
	for _, detail := range details {
		productIDs = append(productIDs, detail.ProductDetailID)
	}

	products, err := service.ProductDetailRepository.FindAllByIds(productIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find product: %w", err)
	}

	// Create maps to store product prices and subtotals
	// productName := make(map[uint]string)
	productPrices := make(map[uint]float64)
	subTotal := make(map[uint]float64)

	for _, product := range products {
		productPrices[product.ID] = product.Price
	}

	for _, detail := range details {
		price := productPrices[detail.ProductDetailID]
		subTotal[detail.ProductDetailID] += float64(detail.Quantity) * price
	}

	return productPrices, subTotal, nil
}

func (service *TransactionImpl) CalculateTotalPrice(details []domain.TransactionDetail) (float64, error) {
	var totalPrice float64

	for _, detail := range details {
		// Fetch the item from the database to get the latest price
		product, err := service.ProductDetailRepository.FindById(detail.ProductDetailID)
		if err != nil {
			return 0, fmt.Errorf("failed to find product: %w", err)
		}

		totalPrice += float64(detail.Quantity) * product.Price
	}	

	return totalPrice, nil
}

func (service *TransactionImpl) CalculateDiscount(id int, totalPrice float64) (float64, error) {
	var discount float64

	point, err := service.ConvertPointRepository.FindById(id)
	if point.Point == 0 {
		return discount , nil
	}
	if err != nil {
		return discount, fmt.Errorf("failed to convert point: %w", err)
	}
	
	calculateDiscount := totalPrice - float64(point.ValuePoint)
	discount = totalPrice - calculateDiscount
	return discount, nil
}


func (service *TransactionImpl) SubtractionPoint(tx *gorm.DB , pointId , membershipId uint) (*domain.Membership, error) {
	
	point, err := service.ConvertPointRepository.FindById(int(pointId))
	if err != nil {
		return nil, fmt.Errorf("failed to find convert point: %w", err)
	}
	
	membership, err := service.MembershipRepository.FindById(int(membershipId))
	if membership.ID == 0 {
		return membership, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find membership: %w", err)
	}
	if membership.Point < point.Point {
		return nil, fmt.Errorf("failed to convert point to discount: %w", err)
	}
	if membership.Point <= 0 {
		return nil, fmt.Errorf("failed to convert point to discount: %w", err)
	}
	membership.Point  = membership.Point - point.Point

	err = service.MembershipRepository.UpdatePoint(tx, membership)

	if err != nil {
		return nil, fmt.Errorf("failed to update point membership: %w", err)
	}
	
	return membership, nil

}

func (service *TransactionImpl) UpdateMemberPoint(tx *gorm.DB, totalPayment float64, membershipID uint) (error) {
    // Hitung jumlah poin berdasarkan total pembayaran
    pointsEarned := uint(totalPayment / 50000) * 5

    if pointsEarned > 0 {
        membership, err := service.MembershipRepository.FindById(int(membershipID))
        if err != nil {
            return fmt.Errorf("failed to find membership: %w", err)
        }
        membership.Point += pointsEarned

        // Simpan perubahan keanggotaan ke database
        err = service.MembershipRepository.UpdatePoint(tx, membership)
        if err != nil {
            return fmt.Errorf("failed to update point membership: %w", err)
        }
    }

    // Ambil data keanggotaan setelah perubahan
    _ , err := service.MembershipRepository.FindById(int(membershipID))
    if err != nil {
        return fmt.Errorf("failed to find updated membership: %w", err)
    }

    return nil
}

func (service *TransactionImpl) CreateInvoice(paymentMethod, paymentType uint) (string, error){
	var method string
	currentTime := time.Now().Unix()
	currentTimeString := strconv.FormatInt(currentTime, 10)
	invoiceNumber := rand.Intn(999) + 1000
	invoiceNumberString := strconv.Itoa(invoiceNumber)
	switch paymentMethod {
    case 1:
        method = "CASH"
    case 2:
        method = "QRIS"
    case 3:
        method = "BANK"
    default:
        return "", fmt.Errorf("unknown payment method")
    }
	invoice := method + currentTimeString + invoiceNumberString
	return invoice, nil
}

func (service *TransactionImpl) MatchingTotalPrice(priceMobile, price float64) error {
	if priceMobile != price {
		return fmt.Errorf("total price does not match")
	}
	return nil
}

func (service *TransactionImpl) MatchingDiscount(discountMobile, discount float64) error {

	if discountMobile != discount {
		return fmt.Errorf("discount does not match")
	}
	return nil
}

func (service *TransactionImpl) MatchingTax(taxMobile, tax float64) error {
	if taxMobile != tax {
		return fmt.Errorf("tax does not match")
	}
	return nil
}

func (service *TransactionImpl) MatchingTotalPayment(totalPaymentMobile, totalPayment float64) error {
	if totalPaymentMobile != totalPayment {
		return fmt.Errorf("total payment does not match")
	}
	return nil
}

func (service *TransactionImpl) NotificationPayment(notificationPayload map[string]interface{}) error {
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return fmt.Errorf("error when get order id : order id not found")
	}

	transactionPaymentStatus, transactionResult, err := service.MidtransCoreApi.CheckTransactionStatus(orderId)
	if err != nil {
		return fmt.Errorf("error when checking transaction status : %s", err.Error())
	}

	err = service.TransactionRepository.UpdateStatusTransactionPayment(transactionPaymentStatus, transactionResult)
	if err != nil {
		return fmt.Errorf("error when update transaction status : %s", err.Error())
	}
	return nil
}

func (service *TransactionImpl) ManualPayment(invoice string) (*domain.Transaction, error) {
	transactionPaymentResult := &domain.PaymentTransactionStatus{
		OrderID: invoice,
		SettlementTime: time.Now(),
		TransactionStatus: "success",
	}
	err := service.TransactionRepository.UpdateStatusTransactionPayment(transactionPaymentResult.TransactionStatus, transactionPaymentResult)
	if err != nil {
		return nil, fmt.Errorf("error when update transaction status : %s", err.Error())
	}

	result, _ := service.TransactionRepository.FindByInvoice(invoice)
	if result == nil {
		return nil, fmt.Errorf("transaction not found")
	}
	return result, nil
}

func (service *TransactionImpl) FindById(id int) (*domain.Transaction, error) {
	existingTransaction, _ := service.TransactionRepository.FindById(id)
	if existingTransaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}

	return existingTransaction, nil
}

func (service *TransactionImpl) FindByInvoice(invoice string) (*domain.Transaction, error) {
	existingTransaction, _ := service.TransactionRepository.FindByInvoice(invoice)
	if existingTransaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}

	return existingTransaction, nil
}

func (service *TransactionImpl) FindByYearly() (*domain.TransactionYearlyRevenue, error) {
	currentYear := time.Now().Year()
	existingTransaction, _ := service.TransactionRepository.FindYearlyRevenue(currentYear)
	if existingTransaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}

	return existingTransaction, nil
}

func (service *TransactionImpl) FindByDaily() (*domain.TransactionDailyRevenue, error) {
	currentDate := time.Now().Format("2006-01-02")
	existingTransaction, err := service.TransactionRepository.FindDailyTransaction(currentDate)
	if err != nil {
		return nil, fmt.Errorf("error when get transaction daily")
	}
	if existingTransaction == nil {
		return nil, fmt.Errorf("transaction daily not found")
	}

	return existingTransaction, nil
}


func (service *TransactionImpl) FindByMonthly() ([]domain.TransactionMonthlyRevenue, error) {
	currentYear := time.Now().Year()
	existingTransaction, _ := service.TransactionRepository.FindMonthlyRevenue(currentYear)
	if existingTransaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}

	return existingTransaction, nil
}

func (service *TransactionImpl) FindAllTransaction() ([]domain.Transaction, int, error) {
	transactions, totalTransaction , err := service.TransactionRepository.FindAllTransaction()
	if err != nil{
		return nil, 0, fmt.Errorf("transaction not found")
	}

	return transactions, totalTransaction, nil
}

func (service *TransactionImpl) FindRecentTransaction() ([]domain.Transaction, error) {
	transactions, err := service.TransactionRepository.FindRecentTransaction()
	if err != nil{
		return nil,fmt.Errorf("transaction not found")
	}

	return transactions,  nil
}

func (service *TransactionImpl)	FindPaginationTransaction(orderBy, QueryLimit, QueryPage string) ([]domain.Transaction, *helpers.Pagination, error) {
	Page, _ := strconv.Atoi(QueryPage)
	Limit, _ := strconv.Atoi(QueryLimit)
	
	Paginate := helpers.Pagination{
		Page: uint(Page),
		Limit: uint(Limit),
	}

	result, paginate, err := service.TransactionRepository.FindPaginationTransaction(orderBy, Paginate)
	if err != nil {
		return nil, nil, fmt.Errorf("pagination transaction error")
	}

	return result, paginate, err
}