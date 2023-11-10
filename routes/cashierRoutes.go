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

func CashierRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	cashierRepository := repository.NewCashierRepository(db)
	cashierService := services.NewCashierService(cashierRepository, validate)
	CashierHandler := handler.NewCashierHandler(cashierService)

	cashierGroup := e.Group("api/v1/cashier")

	cashierGroup.POST("/register", CashierHandler.RegisterCashierHandler)
	cashierGroup.POST("/login", CashierHandler.LoginCashierHandler)

	cashierGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	cashierGroup.GET("/:id", CashierHandler.GetCashierHandler, middleware.AuthMiddleware("Cashier"))
	cashierGroup.GET("", CashierHandler.GetCashiersHandler, middleware.AuthMiddleware("Cashier"))
	cashierGroup.GET("/name/:name", CashierHandler.GetCashierByNameHandler, middleware.AuthMiddleware("Cashier"))
	cashierGroup.PUT("/:id", CashierHandler.UpdateCashierHandler, middleware.AuthMiddleware("Cashier"))
	cashierGroup.DELETE("/:id", CashierHandler.DeleteCashierHandler, middleware.AuthMiddleware("Cashier"))
}