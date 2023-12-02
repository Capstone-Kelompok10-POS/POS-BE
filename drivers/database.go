package drivers

import (
	"fmt"
	"log"
	"os"
	"qbills/models/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("failed connect to database")
	}

	Migrate()

	fmt.Println("Connected to database")

}

func Migrate() {
<<<<<<< Updated upstream
	err := DB.AutoMigrate(
		&schema.SuperAdmin{},
		&schema.Admin{},
<<<<<<< Updated upstream
		&schema.ConvertPoint{},
=======
<<<<<<< Updated upstream
		&schema.ProductType{})
=======
<<<<<<< Updated upstream
    &schema.ConvertPoint{},
=======
<<<<<<< Updated upstream
		&schema.ConvertPoint{},
=======
<<<<<<< Updated upstream
    &schema.ConvertPoint{},
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
    &schema.ConvertPoint{},
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		&schema.ProductType{},
		&schema.Product{},
		&schema.Stock{},
		&schema.Cashier{},
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
		&schema.Membership{})
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
		&schema.Membership{})
=======
		&schema.Membership{},
		&schema.PaymentType{})

=======
<<<<<<< Updated upstream
		&schema.Cashier{},
		&schema.Membership{},
		&schema.ProductType{})
=======
<<<<<<< Updated upstream
    &schema.ConvertPoint{},
>>>>>>> Stashed changes
		&schema.ProductType{},
		&schema.Product{},
		&schema.Stock{},
		&schema.Cashier{},
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		&schema.Membership{},
		&schema.PaymentType{},
		&schema.PaymentMethod{},
		&schema.ProductDetail{})
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes

<<<<<<< Updated upstream
=======
=======
    &schema.Cashier{},
		&schema.ProductType{})
=======
<<<<<<< Updated upstream
	err := DB.AutoMigrate(&schema.SuperAdmin{},
		&schema.Admin{})
=======
	err := DB.AutoMigrate(
		&schema.SuperAdmin{},
		&schema.Admin{},
<<<<<<< Updated upstream
=======
		&schema.Cashier{},
		&schema.Membership{},
<<<<<<< Updated upstream
>>>>>>> Stashed changes
		&schema.ProductType{})
=======
		&schema.PaymentType{},
		&schema.PaymentMethod{},
		&schema.ProductDetail{},
		&schema.Transaction{},
		&schema.TransactionDetail{},
	)

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
	if err != nil {
		log.Fatal("Failed to Migrate Database")
	}
	fmt.Println("Success Migrate Database")
}
