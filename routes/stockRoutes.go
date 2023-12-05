package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
)

func StockRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	stockRepository := repository.NewStockRepository(db)
	productDetailRepository := repository.NewProductDetailRepository(db)
	stockService := services.NewStockService(stockRepository, productDetailRepository, validate)
	stockHandler := handler.NewStockHandler(stockService)

	Group := e.Group("/api/v1/stocks")

	Group.POST("", stockHandler.UpdateStockHandler)
	Group.GET("", stockHandler.FindAllStockHandler)
	Group.GET("/:id", stockHandler.FindByIdStockHandler)
	Group.GET("/get/increase", stockHandler.FindIncreaseStockHandler)
	Group.GET("/get/decrease", stockHandler.FindDecreaseStockHandler)

}
