package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"
	res "qbills/utils/response"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransaction(ctx echo.Context, request web.TransactionCreateRequest) (*web.TransactionResponse, error)
	FindById(ctx echo.Context, id int) (*domain.Transaction, error)
	GetPricesAndSubTotal(details []domain.TransactionDetail) (map[uint]float64, map[uint]float64, error)
	CalculateTotalPrice(details []domain.TransactionDetail) (float64, error)
	CalculateDiscount(id int, totalPrice float64) (float64, error)
	SubtractionPoint(tx *gorm.DB , pointId , membershipId uint) (*domain.Membership, error)
}

type TransactionImpl struct {
	TransactionRepository repository.TransactionRepository
	ProductDetailRepository repository.ProductDetailRepository
	ConvertPointRepository repository.ConvertPointRepository
	MembershipRepository repository.MembershipRepository
	validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductDetailRepository, convertPointRepository repository.ConvertPointRepository, membershipRepository repository.MembershipRepository ,validate *validator.Validate) *TransactionImpl {
	return &TransactionImpl{
		TransactionRepository: transactionRepository,
		ProductDetailRepository: productRepository,
		ConvertPointRepository: convertPointRepository,
		MembershipRepository: membershipRepository,
		validate:              validate,
	}
}

func (service *TransactionImpl) CreateTransaction(ctx echo.Context, request web.TransactionCreateRequest) (*web.TransactionResponse, error) {
	tx := service.TransactionRepository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
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
	tax := totalPrice - discount
	totalTax := (0.10 * tax)

	//Add status Transaction
	status := "PENDING"

	//calculate total payment 
	totalPayment := (totalPrice - discount) + totalTax

	transaction := req.TransactionCreateRequestToTransactionDomain(request, 
		web.TransactionCreate{
			Status: status,
			Discount: discount,
			TotalPrice: totalPrice,
			Tax: tax,
			TotalPayment: totalPayment,
		}, web.TransactionDetailCreate{
			ProductPrice : productPrice,
			SubTotal: subTotal,
		})
	
	result, err := service.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction %w", err)
	}

	_ , err = service.SubtractionPoint(tx, result.ConvertPointID, result.MembershipID)
	if err != nil {
		return nil, fmt.Errorf("error when decreasing point membership %w", err)
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
	
	response := res.TransactionDomainToTransactionResponse(result)
	return response, nil
}

func (service *TransactionImpl) FindById(ctx echo.Context, id int) (*domain.Transaction, error) {
	existingTransaction, _ := service.TransactionRepository.FindById(id)
	if existingTransaction == nil {
		return nil, fmt.Errorf("product type not found")
	}

	return existingTransaction, nil
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