package repository

import (
<<<<<<< Updated upstream
	"fmt"
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
	"gorm.io/gorm"
	"os"
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/utils/helpers"
	req "qbills/utils/request"
	res "qbills/utils/response"
<<<<<<< Updated upstream
	"strconv"
=======
=======
	"fmt"
<<<<<<< Updated upstream
	"gorm.io/gorm"
=======
>>>>>>> Stashed changes
	"os"
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/utils/helpers"
	req "qbills/utils/request"
	res "qbills/utils/response"
	"strconv"
<<<<<<< Updated upstream
=======

	"gorm.io/gorm"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
)

type ProductRepository interface {
	Create(request *domain.Product) (*domain.Product, error)
	Update(request *domain.Product, id uint) (*domain.Product, error)
	FindById(id uint) (*domain.Product, error)
	FindAll() ([]domain.Product, error)
	FindByName(name string) ([]domain.Product, error)
<<<<<<< Updated upstream
	FindByCategory(ProductTypeID uint) ([]domain.Product, error)
	Delete(id uint) error
	FindPaginationProduct(orderBy string, paginate helpers.Pagination) ([]domain.Product, *helpers.Pagination, error)
=======
<<<<<<< Updated upstream
	Delete(id uint) error
=======
	FindByCategory(ProductTypeID uint) ([]domain.Product, error)
	Delete(id uint) error
	FindPaginationProduct(orderBy string, paginate helpers.Pagination) ([]domain.Product, *helpers.Pagination, error)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}

func (repository *ProductRepositoryImpl) Create(request *domain.Product) (*domain.Product, error) {
	ProductDB := req.ProductDomainToProductSchema(*request)

	result := repository.DB.Create(&ProductDB)

	if result.Error != nil {
		return nil, result.Error
	}

	response := res.ProductSchemaToProductDomain(ProductDB)

	return response, nil
}

func (repository *ProductRepositoryImpl) Update(request *domain.Product, id uint) (*domain.Product, error) {
<<<<<<< Updated upstream
	result := repository.DB.Table("products").Where("id = ?", id).Updates(request)
=======
<<<<<<< Updated upstream
	result := repository.DB.Table("products").Where("id = ?", id).Updates(domain.Product{
		ProductTypeID: request.ProductTypeID,
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		Stock:         request.Stock,
		Size:          request.Size,
		Image:         request.Image,
	})
=======
<<<<<<< Updated upstream
	result := repository.DB.Table("products").Where("id = ?", id).Save(request)
=======
<<<<<<< Updated upstream
	result := repository.DB.Table("products").Where("id = ?", id).Updates(domain.Product{
		ProductTypeID: request.ProductTypeID,
		Name:          request.Name,
		Ingredients:   request.Ingredients,
		Price:         request.Price,
		TotalStock:    request.TotalStock,
		Size:          request.Size,
		Image:         request.Image,
	})
=======
<<<<<<< Updated upstream
	result := repository.DB.Table("products").Where("id = ?", id).Save(request)
=======
	result := repository.DB.Table("products").Where("id = ?", id).Updates(request)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (repository *ProductRepositoryImpl) FindById(id uint) (*domain.Product, error) {
	product := domain.Product{}

<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL").First(&product, id)
=======
<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL").First(&product, id)
=======
	result := repository.DB.Preload("ProductType").Preload("Admin").Preload("ProductDetail").Where("deleted_at IS NULL").First(&product, id)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *ProductRepositoryImpl) FindAll() ([]domain.Product, error) {
	product := []domain.Product{}

<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL").Find(&product)
=======
<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL").Find(&product)
=======
	result := repository.DB.Preload("ProductType").Preload("Admin").Preload("ProductDetail").Where("deleted_at IS NULL").Find(&product)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

<<<<<<< Updated upstream
func (repository *ProductRepositoryImpl) FindByCategory(ProductTypeID uint) ([]domain.Product, error) {
	products := []domain.Product{}

	result := repository.DB.Preload("ProductType").Preload("Admin").Where("product_type_id = ?", ProductTypeID).Find(&products)
=======
<<<<<<< Updated upstream
=======
func (repository *ProductRepositoryImpl) FindByCategory(ProductTypeID uint) ([]domain.Product, error) {
	products := []domain.Product{}

<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("product_type_id = ?", ProductTypeID).Find(&products)
=======
	result := repository.DB.Preload("ProductType").Preload("Admin").Preload("ProductDetail").Where("product_type_id = ?", ProductTypeID).Find(&products)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
func (repository *ProductRepositoryImpl) FindByName(name string) ([]domain.Product, error) {
	products := []domain.Product{}

	// Menambahkan klausa pencarian berdasarkan nama ke query
<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL AND name LIKE ?", "%"+name+"%").Find(&products)
=======
<<<<<<< Updated upstream
	result := repository.DB.Preload("ProductType").Preload("Admin").Where("deleted_at IS NULL AND name LIKE ?", "%"+name+"%").Find(&products)
=======
	result := repository.DB.Preload("ProductType").Preload("Admin").Preload("ProductDetail").Where("deleted_at IS NULL AND name LIKE ?", "%"+name+"%").Find(&products)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	// Memeriksa kesalahan pada query
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) Delete(id uint) error {
	result := repository.DB.Delete(&schema.Product{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes

func (repository *ProductRepositoryImpl) FindPaginationProduct(orderBy string, paginate helpers.Pagination) ([]domain.Product, *helpers.Pagination, error) {
	var products []domain.Product

<<<<<<< Updated upstream
	result := repository.DB.Scopes(helpers.Paginate(products, &paginate, repository.DB)).Preload("Admin").Preload("ProductType")
=======
<<<<<<< Updated upstream
	result := repository.DB.Scopes(helpers.Paginate(products, &paginate, repository.DB)).Preload("Admin").Preload("ProductType")
=======
	result := repository.DB.Scopes(helpers.Paginate(products, &paginate, repository.DB)).Preload("Admin").Preload("ProductType").Preload("ProductDetail")
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if orderBy != "" {
		result.Order("name " + orderBy).Find(&products)
	} else {
		result.Find(&products)
	}

	if result.Error != nil {
		return nil, nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil, fmt.Errorf("products is empty")
	}

	if paginate.Page <= 1 {
		paginate.PreviousPage = ""
	} else {
		paginate.PreviousPage = os.Getenv("MAIN_URL") + "/api/" + os.Getenv("API_VERSION") + "/product/pagination?page=" + strconv.Itoa(int(paginate.Page)-1)
	}

	if paginate.Page >= paginate.TotalPage {
		paginate.NextPage = ""
	} else {
		paginate.NextPage = os.Getenv("MAIN_URL") + "/api/" + os.Getenv("API_VERSION") + "/product/pagination?page=" + strconv.Itoa(int(paginate.Page)+1)
	}

	return products, &paginate, nil
}
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
