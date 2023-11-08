package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CashierRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	cashierRepository := repository.NewCashierRepository(db)
	cashierService := services.NewCashierService(cashierRepository, validate)
	CashierHandler := handler.NewCashierHandler(cashierService)

	cashierGroup := e.Group("api/v1/cashier")

	cashierGroup.POST("/login", CashierHandler.LoginCashierHandler)

	cashierGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
}