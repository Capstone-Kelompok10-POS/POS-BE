package drivers

import (
	"fmt"
	"qbills/configs"
	"qbills/models/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func NewMySQLConnection(config *configs.MySQLConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	Migrate(db)

	return db, nil

}

<<<<<<< Updated upstream
func Migrate() {
	err := DB.AutoMigrate(
=======
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
>>>>>>> Stashed changes
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
		&schema.ProductType{})
=======
		&schema.ConvertPoint{},
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		&schema.ProductType{},
		&schema.Product{},
		&schema.Stock{},
		&schema.Cashier{},
<<<<<<< Updated upstream
		&schema.Membership{},
		&schema.PaymentType{},
		&schema.PaymentMethod{},
		&schema.ProductDetail{})

=======
<<<<<<< Updated upstream
		&schema.Membership{})

=======
		&schema.Membership{},
		&schema.PaymentType{},
<<<<<<< Updated upstream
		&schema.PaymentMethod{},
		&schema.ProductDetail{})
=======
<<<<<<< Updated upstream
		&schema.PaymentMethod{})
=======
<<<<<<< Updated upstream
		&schema.PaymentMethod{},
		&schema.ProductDetail{})
=======
<<<<<<< Updated upstream
		&schema.PaymentMethod{})
=======
		&schema.PaymentMethod{},
<<<<<<< Updated upstream
		&schema.ProductDetail{})
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
=======
		&schema.ProductDetail{},
		&schema.Transaction{},
		&schema.TransactionDetail{},
		&schema.TransactionPayment{},
	)
}
>>>>>>> Stashed changes
