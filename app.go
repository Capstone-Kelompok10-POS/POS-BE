package main

import (
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"qbills/drivers"
	"qbills/routes"
<<<<<<< Updated upstream
	"qbills/utils/helpers"
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	"qbills/utils/helpers"
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	"qbills/utils/helpers"
=======
<<<<<<< Updated upstream
=======
	"qbills/utils/helpers"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
)

func main() {
	myApp := echo.New()
	validate := validator.New()
	helpers.ConnectAWS()

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	drivers.ConnectDB()
	drivers.Migrate()

	myApp.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Q Bills API Services")
	})

	routes.AdminRoutes(myApp, drivers.DB, validate)
	routes.CashierRoutes(myApp, drivers.DB, validate)
	routes.SuperAdminRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
	routes.ConvertPointRoutes(myApp, drivers.DB, validate)
	routes.ProductTypeRoutes(myApp, drivers.DB, validate)
	routes.ProductRoutes(myApp, drivers.DB, validate)
	routes.StockRoutes(myApp, drivers.DB, validate)
	routes.MembershipRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
	routes.ProductTypeRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
	routes.ConvertPointRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	routes.ConvertPointRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	routes.ConvertPointRoutes(myApp, drivers.DB, validate)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	routes.ProductTypeRoutes(myApp, drivers.DB, validate)
	routes.ProductRoutes(myApp, drivers.DB, validate)
	routes.StockRoutes(myApp, drivers.DB, validate)
	routes.MembershipRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	routes.PaymentTypeRoutes(myApp, drivers.DB, validate)
	routes.PaymentMethodRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
	routes.ProductDetailRoutes(myApp, drivers.DB, validate)
=======
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	routes.PaymentTypeRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
=======
	routes.PaymentMethodRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
=======
	routes.MembershipCardRoutes(myApp, drivers.DB)
	routes.PaymentTypeRoutes(myApp, drivers.DB, validate)
	routes.PaymentMethodRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
	routes.ProductTypeRoutes(myApp, drivers.DB, validate)
=======
<<<<<<< Updated upstream
=======
	routes.ProductTypeRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
=======
	routes.MembershipRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
	routes.MembershipCardRoutes(myApp, drivers.DB)
	routes.PaymentTypeRoutes(myApp, drivers.DB, validate)
	routes.PaymentMethodRoutes(myApp, drivers.DB, validate)
	routes.ProductDetailRoutes(myApp, drivers.DB, validate)
<<<<<<< Updated upstream
=======
	routes.TransactionRoutes(myApp, drivers.DB, validate)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	myApp.Pre(middleware.RemoveTrailingSlash())
	myApp.Use(middleware.CORS())
	myApp.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	myApp.Logger.Fatal(myApp.Start(":8005"))
}
