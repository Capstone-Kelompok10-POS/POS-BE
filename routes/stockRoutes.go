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

func StockRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	stockRepository := repository.NewStockRepository(db)
	productRepository := repository.NewProductRepository(db)
	stockService := services.NewStockService(stockRepository, productRepository, validate)
	stockHandler := handler.NewStockHandler(stockService)

	StockGroup := e.Group("/api/v1/product/stocks")

	StockGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	StockGroup.POST("", stockHandler.UpdateStockHandler)
	StockGroup.GET("", stockHandler.FindAllStockHandler)
	StockGroup.GET("/:id", stockHandler.FindByIdStockHandler)
	StockGroup.GET("/get/increase", stockHandler.FindIncreaseStockHandler)
	StockGroup.GET("/get/decrease", stockHandler.FindDecreaseStockHandler)

}
