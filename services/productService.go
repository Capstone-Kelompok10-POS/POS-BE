package services

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"
	"strconv"
	"strings"
)

type ProductService interface {
	CreateProductService(ctx echo.Context, request web.ProductCreateRequest) (*domain.Product, error)
	UpdateProductService(ctx echo.Context, request web.ProductUpdateRequest, id uint) (*domain.Product, error)
	FindByIdProductService(ctx echo.Context, id uint) (*domain.Product, error)
	FindByNameProductService(ctx echo.Context, name string) ([]domain.Product, error)
	FindAllProductService(ctx echo.Context) ([]domain.Product, error)
	FindByCategoryProductService(ctx echo.Context, productTypeID uint) ([]domain.Product, error)
	DeleteProductService(ctx echo.Context, id uint) error
	UploadImageProduct(ctx echo.Context) (string, error)
	FindPaginationProduct(ctx echo.Context) ([]domain.Product, *helpers.Pagination, error)
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		validate:          validate,
	}
}

func (service *ProductServiceImpl) CreateProductService(ctx echo.Context, request web.ProductCreateRequest) (*domain.Product, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	product := req.ProductCreateRequestToProductDomain(request)

	result, err := service.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *ProductServiceImpl) UpdateProductService(ctx echo.Context, request web.ProductUpdateRequest, id uint) (*domain.Product, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	exitingAdmin, _ := service.ProductRepository.FindById(id)

	if exitingAdmin == nil {
		return nil, fmt.Errorf("product not found")
	}

	product := req.ProductUpdateRequestToProductDomain(request)

	result, err := service.ProductRepository.Update(product, id)

	if err != nil {
		return nil, fmt.Errorf("error when updating data product: %s", err.Error())
	}

	return result, nil

}

func (service *ProductServiceImpl) FindByIdProductService(ctx echo.Context, id uint) (*domain.Product, error) {
	result, err := service.ProductRepository.FindById(id)

	if err != nil {
		return nil, fmt.Errorf("product not found")
	}

	return result, nil
}

func (service *ProductServiceImpl) FindAllProductService(ctx echo.Context) ([]domain.Product, error) {
	product, err := service.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) FindByNameProductService(ctx echo.Context, name string) ([]domain.Product, error) {
	products, err := service.ProductRepository.FindByName(name)

	if err != nil {
		return nil, fmt.Errorf("failed to find products with the name %s: %s", name, err.Error())
	}

	return products, nil
}

func (service *ProductServiceImpl) FindByCategoryProductService(ctx echo.Context, productTypeID uint) ([]domain.Product, error) {
	product, err := service.ProductRepository.FindByCategory(productTypeID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) DeleteProductService(ctx echo.Context, id uint) error {
	exitingProduct, _ := service.ProductRepository.FindById(id)
	if exitingProduct == nil {
		return fmt.Errorf("product not found")
	}

	err := service.ProductRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (service *ProductServiceImpl) UploadImageProduct(ctx echo.Context) (string, error) {

	// Set the path to your service account JSON file
	serviceAccountKeyPath := "credentials.json"

	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error initializing app: %v", err))
	}

	// Generate a new UUID
	newUUID := uuid.New()

	// Convert the UUID to a string
	uuidString := newUUID.String()

	// Set the destination path in Firebase Storage
	storagePath := "product/" + uuidString + ".png"

	// Open the uploaded file
	file, err := ctx.FormFile("image")
	if err != nil {
		return "", ctx.String(http.StatusBadRequest, fmt.Sprintf("Error reading uploaded file: %v", err))
	}

	src, err := file.Open()
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error opening uploaded file: %v", err))
	}
	defer src.Close()

	// Initialize Google Cloud Storage client
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error creating storage client: %v", err))
	}
	defer client.Close()

	// Specify the name of your Firebase Storage bucket
	bucketName := "qbils-d46b3.appspot.com"

	// Set the appropriate MIME type based on the file extension
	fileExtension := strings.TrimLeft(filepath.Ext(file.Filename), ".")
	var contentType string
	switch fileExtension {
	case "jpg", "jpeg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	default:
		return "", ctx.String(http.StatusBadRequest, fmt.Sprintf("Unsupported file format: %s", fileExtension))
	}

	// Upload the file to Firebase Storage with the determined content type
	object := client.Bucket(bucketName).Object(storagePath)
	wc := object.NewWriter(context.Background())
	wc.ContentType = contentType

	if _, err := io.Copy(wc, src); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error copying file to Firebase Storage: %v", err))
	}
	if err := wc.Close(); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error closing writer: %v", err))
	}

	// Set ACL for public read access after creating the object
	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error setting ACL: %v", err))
	}

	// Get the download URL
	_, err = object.Attrs(context.Background())
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file attributes: %v", err))
	}

	// Return the read-only URL to the client
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucketName, url.QueryEscape(storagePath))

	return url, nil
}

func (service *ProductServiceImpl) FindPaginationProduct(ctx echo.Context) ([]domain.Product, *helpers.Pagination, error) {

	orderBy := ctx.QueryParam("orderBy")
	QueryLimit := ctx.QueryParam("limit")
	QueryPage := ctx.QueryParam("page")

	Page, _ := strconv.Atoi(QueryPage)
	Limit, _ := strconv.Atoi(QueryLimit)

	Paginate := helpers.Pagination{
		Page:  uint(Page),
		Limit: uint(Limit),
	}

	result, paginate, err := service.ProductRepository.FindPaginationProduct(orderBy, Paginate)
	if err != nil {
		return nil, nil, fmt.Errorf("Product is empty")
	}

	return result, paginate, nil
}
