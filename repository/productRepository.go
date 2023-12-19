package repository

import (
	"fmt"
	"os"
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/utils/helpers"
	"strconv"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(request *domain.Product) (*domain.Product, error)
	Update(request *domain.Product, id uint) (*domain.Product, error)
	FindById(id uint) (*domain.Product, error)
	FindAll() ([]domain.Product, int, error)
	FindByName(name string) ([]domain.Product, error)
	FindByCategory(ProductTypeID uint) ([]domain.Product, error)
	Delete(id uint) error
	FindPaginationProduct(orderBy string, paginate helpers.Pagination) ([]domain.Product, *helpers.Pagination, error)
	FindBestSellingProduct() ([]domain.BestSellingProduct, error)
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}

func (repository *ProductRepositoryImpl) Create(request *domain.Product) (*domain.Product, error) {

	result := repository.DB.Create(&request)

	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (repository *ProductRepositoryImpl) Update(request *domain.Product, id uint) (*domain.Product, error) {
	result := repository.DB.Table("products").Where("id = ?", id).Updates(request)

	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (repository *ProductRepositoryImpl) FindById(id uint) (*domain.Product, error) {
	product := domain.Product{}

	result := repository.DB.Preload("ProductType").Preload("Admin").Preload("ProductDetail", "deleted_at IS NULL").Where("deleted_at IS NULL").First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *ProductRepositoryImpl) FindAll() ([]domain.Product, int, error) {
	products := []domain.Product{}

	result := repository.DB.
		Preload("ProductType").
		Preload("ProductDetail", "deleted_at IS NULL").
		Where("deleted_at IS NULL").
		Find(&products)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	totalProducts := len(products)
	return products, totalProducts, nil
}

func (repository *ProductRepositoryImpl) FindByCategory(ProductTypeID uint) ([]domain.Product, error) {
	products := []domain.Product{}

	result := repository.DB.
		Preload("ProductType").
		Preload("Admin").
		Preload("ProductDetail", "deleted_at IS NULL").
		Where("deleted_at IS NULL AND product_type_id = ?", ProductTypeID).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) FindByName(name string) ([]domain.Product, error) {
	products := []domain.Product{}

	// Menambahkan klausa pencarian berdasarkan nama ke query
	result := repository.DB.
		Preload("ProductType").
		Preload("Admin").
		Preload("ProductDetail", "deleted_at IS NULL").
		Where("deleted_at IS NULL AND name LIKE ?", "%"+name+"%").
		Find(&products)

	// Memeriksa kesalahan pada query
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) Delete(id uint) error {
	result := repository.DB.Where("deleted_at IS NULL AND id = ?", id).Delete(&schema.Product{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *ProductRepositoryImpl) FindPaginationProduct(orderBy string, paginate helpers.Pagination) ([]domain.Product, *helpers.Pagination, error) {
	var products []domain.Product

	result := repository.DB.Scopes(helpers.Paginate(products, &paginate, repository.DB)).Preload("Admin").Preload("ProductType").Preload("ProductDetail", "deleted_at IS NULL")

	if orderBy != "" {
		result.Order("name " + orderBy).Where("products.deleted_at IS NULL").Find(&products)
	} else {
		result.Where("products.deleted_at IS NULL").Find(&products)
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

func (repository *ProductRepositoryImpl) FindBestSellingProduct() ([]domain.BestSellingProduct, error) {
	bestProduct := []domain.BestSellingProduct{}

	query := `SELECT
    p.id as product_id,
    p.name as product_name,
    pd.size AS product_size,
    p.image AS product_image,
    pt.type_name as product_type_name,
    pd.price AS product_price,
    SUM(td.quantity) as total_quantity,
    SUM(td.sub_total) as amount
	FROM
		transaction_details td
	JOIN
		product_details pd ON td.product_detail_id = pd.id
	JOIN
		transactions t ON td.transaction_id = t.id
	JOIN
		transaction_payments tp ON t.id = tp.transaction_id
	JOIN
		products p ON pd.product_id = p.id
	LEFT JOIN
		product_types pt ON p.product_type_id = pt.id
	WHERE
		tp.payment_status = 'success'
	GROUP BY
		p.id, p.name, p.image, pd.price, pt.type_name, pd.size
	ORDER BY
		total_quantity DESC
	LIMIT 3;
	`
	result := repository.DB.Raw(query).Preload("ProductDetail", "deleted_at IS NULL").Scan(&bestProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return bestProduct,  nil
}
