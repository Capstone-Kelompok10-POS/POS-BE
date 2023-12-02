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

func PaymentMethodRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	paymentMethodRepository := repository.NewPaymentMethodRepository(db)
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepository, validate)
	paymentMethodHandler := handler.NewPaymentMethodHandler(paymentMethodService)

	paymentMethodGroup := e.Group("api/v1/payment/method")

	paymentMethodGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	paymentMethodGroup.POST("", paymentMethodHandler.CreatePaymentMethodHandler)
	paymentMethodGroup.GET("/:id", paymentMethodHandler.GetPaymentMethodHandler)
	paymentMethodGroup.GET("", paymentMethodHandler.GetPaymentMethodsHandler)
	paymentMethodGroup.GET("/name/:name", paymentMethodHandler.GetPaymentMethodByNameHandler)
	paymentMethodGroup.PUT("/:id", paymentMethodHandler.UpdatePaymentMethodHandler)
	paymentMethodGroup.DELETE("/:id", paymentMethodHandler.DeletePaymentMethodHandler)
}
