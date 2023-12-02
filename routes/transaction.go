package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	transactionRepository := repository.NewTransactionRepository(db)
	productDetailRepository := repository.NewProductDetailRepository(db)
	convertPointRepository := repository.NewConvertPointRepository(db)
	membershipRepository := repository.NewMembershipRepository(db)
	transactionService := services.NewTransactionService(transactionRepository, productDetailRepository, convertPointRepository,membershipRepository,validate)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	transactionGroup := e.Group("/api/v1/transaction")
	transactionGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	transactionGroup.POST("", transactionHandler.CreateTransactionHandler, middleware.AuthMiddleware("Cashier"))
	transactionGroup.GET("/:id", transactionHandler.GetTransactionHandler)
}
