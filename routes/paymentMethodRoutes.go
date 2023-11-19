package routes

import (
	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
)

func PaymentMethodRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	paymentMethodRepository := repository.NewPaymentTypeRepository(db)
	paymentMethodService := services.NewPaymentTypeService(paymentMethodRepository, validate)
	paymentMethodHandler := handler.NewPaymentTypeHandler(paymentMethodService)

	paymentMethodGroup := e.Group("api/v1/paymentMethod")

	paymentMethodGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	paymentMethodGroup.POST("/register", paymentMethodHandler.CreatePaymentTypeHandler)
	paymentMethodGroup.GET("/:id", paymentMethodHandler.UpdatePaymentTypeHandler)
	paymentMethodGroup.GET("", paymentMethodHandler.GetPaymentTypesHandler)
	paymentMethodGroup.GET("/name/:name", paymentMethodHandler.GetPaymentTypeByNameHandler)
	paymentMethodGroup.PUT("/:id", paymentMethodHandler.UpdatePaymentTypeHandler)
	paymentMethodGroup.DELETE("/:id", paymentMethodHandler.DeletePaymentTypeHandler)
}
