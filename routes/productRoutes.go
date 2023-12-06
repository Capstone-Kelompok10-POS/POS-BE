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

func ProductRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productRepository := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepository, validate)
	ProductHandler := handler.NewProductHandler(productService)

	productGroup := e.Group("api/v1/product")
	productGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	productGroup.POST("", ProductHandler.CreateProductHandler)
	productGroup.GET("", ProductHandler.FindPaginationProduct)
	productGroup.GET("/:id", ProductHandler.GetProductHandler)
	productGroup.GET("/all", ProductHandler.GetProductsHandler)
	productGroup.GET("/search/:name", ProductHandler.GetProductByNameHandler)
	productGroup.GET("/category/:productTypeID", ProductHandler.GetProductByCategoryHandler)
	productGroup.PUT("/:id", ProductHandler.UpdateProductHandler)
	productGroup.DELETE("/:id", ProductHandler.DeleteProductHandler)
<<<<<<< Updated upstream
}
=======
}
>>>>>>> Stashed changes
