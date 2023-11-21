package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
)

func PaymentTypeRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	paymentTypeRepository := repository.NewPaymentTypeRepository(db)
	paymentTypeService := services.NewPaymentTypeService(paymentTypeRepository, validate)
	paymentTypeHandler := handler.NewPaymentTypeHandler(paymentTypeService)

	paymentTypeGroup := e.Group("api/v1/paymentType")

	//paymentTypeGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	paymentTypeGroup.POST("", paymentTypeHandler.CreatePaymentTypeHandler)
	paymentTypeGroup.GET("/:id", paymentTypeHandler.GetPaymentTypeHandler)
	paymentTypeGroup.GET("", paymentTypeHandler.GetPaymentTypesHandler)
	paymentTypeGroup.GET("/name/:name", paymentTypeHandler.GetPaymentTypeByNameHandler)
	paymentTypeGroup.PUT("/:id", paymentTypeHandler.UpdatePaymentTypeHandler)
	paymentTypeGroup.DELETE("/:id", paymentTypeHandler.DeletePaymentTypeHandler)
	paymentTypeGroup.GET("/upload", paymentTypeHandler.UploadBarcode)
}
