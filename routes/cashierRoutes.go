package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"
	"qbills/utils/helpers/password"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CashierRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate, password password.PasswordHandler) {
	cashierRepository := repository.NewCashierRepository(db)
	cashierService := services.NewCashierService(cashierRepository, validate, password)
	CashierHandler := handler.NewCashierHandler(cashierService)

	cashierGroup := e.Group("api/v1/cashier")

	cashierGroup.POST("/login", CashierHandler.LoginCashierHandler)

	cashierGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	cashierGroup.POST("/register", CashierHandler.RegisterCashierHandler, middleware.AuthMiddleware("Admin"))
	cashierGroup.GET("/:id", CashierHandler.GetCashierHandler, middleware.AuthMiddleware("Admin"))
	cashierGroup.GET("", CashierHandler.GetCashiersHandler, middleware.AuthMiddleware("Admin"))
	cashierGroup.GET("/username/:username", CashierHandler.GetCashierByUsernameHandler, middleware.AuthMiddleware("Admin"))
	cashierGroup.PUT("/:id", CashierHandler.UpdateCashierHandler, middleware.AuthMiddleware("Admin"))
	cashierGroup.DELETE("/:id", CashierHandler.DeleteCashierHandler, middleware.AuthMiddleware("Admin"))
}
