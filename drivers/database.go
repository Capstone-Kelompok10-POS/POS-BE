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

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&schema.SuperAdmin{},
		&schema.Admin{},
		&schema.Cashier{},
		&schema.Membership{},
		&schema.ConvertPoint{},
		&schema.ProductType{},
		&schema.Product{},
		&schema.Stock{},
		&schema.PaymentType{},
		&schema.PaymentMethod{},
		&schema.ProductDetail{},
		&schema.Transaction{},
		&schema.TransactionDetail{},
    &schema.TransactionPayment{},
	)
	if err != nil {
		log.Fatal("Failed to Migrate Database")
	}
	fmt.Println("Success Migrate Database")
}
