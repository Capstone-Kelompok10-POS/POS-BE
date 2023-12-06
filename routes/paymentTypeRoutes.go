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

func PaymentTypeRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	paymentTypeRepository := repository.NewPaymentTypeRepository(db)
	paymentTypeService := services.NewPaymentTypeService(paymentTypeRepository, validate)
	paymentTypeHandler := handler.NewPaymentTypeHandler(paymentTypeService)

	paymentTypeGroup := e.Group("api/v1/payment/type")

	paymentTypeGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	paymentTypeGroup.POST("", paymentTypeHandler.CreatePaymentTypeHandler)
	paymentTypeGroup.GET("/:id", paymentTypeHandler.GetPaymentTypeHandler)
	paymentTypeGroup.GET("", paymentTypeHandler.GetPaymentTypesHandler)
	paymentTypeGroup.GET("/name/:name", paymentTypeHandler.GetPaymentTypeByNameHandler)
	paymentTypeGroup.PUT("/:id", paymentTypeHandler.UpdatePaymentTypeHandler)
	paymentTypeGroup.DELETE("/:id", paymentTypeHandler.DeletePaymentTypeHandler)
}
