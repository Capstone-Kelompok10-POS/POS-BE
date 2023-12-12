package main

import (
	"net/http"

	"qbills/configs"
	"qbills/drivers"
	"qbills/routes"
	"qbills/utils/helpers/midtrans"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	myApp := echo.New()
	validate := validator.New()

	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config:", err.Error())
	}

	db, err := drivers.NewMySQLConnection(&config.MySQL)
	if err != nil {
		logrus.Fatal("Error connecting to MySQL:", err.Error())
	}

	midtransCoreApi := midtrans.NewMidtransCoreApi(&config.Midtrans)

	myApp.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Q Bills API Services")
	})

	routes.AdminRoutes(myApp, db, validate)
	routes.CashierRoutes(myApp, db, validate)
	routes.SuperAdminRoutes(myApp, db, validate)
	routes.ProductRoutes(myApp, db, validate)
	routes.StockRoutes(myApp, db, validate)
	routes.ConvertPointRoutes(myApp, db, validate)
	routes.ProductTypeRoutes(myApp, db, validate)
	routes.MembershipRoutes(myApp, db, validate)
	routes.MembershipCardRoutes(myApp, db)
	routes.PaymentTypeRoutes(myApp, db, validate)
	routes.PaymentMethodRoutes(myApp, db, validate)
	routes.ProductDetailRoutes(myApp, db, validate)
	routes.MemberShipPointRoutes(myApp, db, validate)
	routes.TransactionRoutes(myApp, db, midtransCoreApi, validate)

	myApp.Pre(middleware.RemoveTrailingSlash())
	myApp.Use(middleware.CORS())
	myApp.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	myApp.Logger.Fatal(myApp.Start(":8080"))
}
