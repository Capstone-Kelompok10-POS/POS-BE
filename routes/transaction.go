package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"
	"qbills/utils/helpers/midtrans"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoutes(e *echo.Echo, db *gorm.DB, midtransCoreApi midtrans.MidtransCoreApi,validate *validator.Validate) {
	transactionRepository := repository.NewTransactionRepository(db)
	productDetailRepository := repository.NewProductDetailRepository(db)
	convertPointRepository := repository.NewConvertPointRepository(db)
	membershipRepository := repository.NewMembershipRepository(db)
	PaymentMethodRepository := repository.NewPaymentMethodRepository(db)
	transactionService := services.NewTransactionService(transactionRepository, productDetailRepository, convertPointRepository,membershipRepository,PaymentMethodRepository, midtransCoreApi,validate)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	transactionGroup := e.Group("/api/v1/transaction")
	transactionGroup.POST("/notifications", transactionHandler.NotificationPayment)
	
	transactionGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	transactionGroup.POST("", transactionHandler.CreateTransactionHandler, middleware.AuthMiddleware("Cashier"))
	transactionGroup.GET("/:id", transactionHandler.GetTransactionHandler)
	transactionGroup.GET("s", transactionHandler.GetTransactionsHandler)
	transactionGroup.GET("s/recent", transactionHandler.GetRecentTransactionsHandler)
	transactionGroup.GET("/revenue/yearly", transactionHandler.GetTransactionYearlyHandler, middleware.AuthMiddleware("Admin"))
	transactionGroup.GET("s/revenue/monthly", transactionHandler.GetTransactionMonthlyHandler, middleware.AuthMiddleware("Admin"))
	transactionGroup.GET("s/pagination", transactionHandler.FindPaginationTransaction)
	transactionGroup.PUT("/payment", transactionHandler.UpdateStatusTransactionPaymentHandler, middleware.AuthMiddleware("Cashier"))
	// transactionGroup.GET("", transactionHandler.GetTransactionsHandler)
}
