package handler

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	"qbills/utils/request"
	res "qbills/utils/response"
	"strconv"
	"strings"
	"time"
)

type ProductHandler interface {
	CreateProductHandler(ctx echo.Context) error
	UpdateProductHandler(ctx echo.Context) error
	GetProductHandler(ctx echo.Context) error
	GetProductsHandler(ctx echo.Context) error
	GetProductByNameHandler(ctx echo.Context) error
	DeleteProductHandler(ctx echo.Context) error
}

type ProductHandlerImpl struct {
	ProductService services.ProductService
}

func NewProductHandler(ProductService services.ProductService) ProductHandler {
	return &ProductHandlerImpl{ProductService: ProductService}
}

func (c *ProductHandlerImpl) CreateProductHandler(ctx echo.Context) error {

	// Set the path to your service account JSON file
	serviceAccountKeyPath := "credentials.json"

	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error initializing app: %v", err))
	}

	// Set the destination path in Firebase Storage
	storagePath := "product/" + time.Now().Format("2006-01-02_15:04:05") + ".png"

	// Open the uploaded file
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Error reading uploaded file: %v", err))
	}

	src, err := file.Open()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error opening uploaded file: %v", err))
	}
	defer src.Close()

	// Initialize Google Cloud Storage client
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error creating storage client: %v", err))
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
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unsupported file format: %s", fileExtension))
	}

	// Upload the file to Firebase Storage with the determined content type
	object := client.Bucket(bucketName).Object(storagePath)
	wc := object.NewWriter(context.Background())
	wc.ContentType = contentType

	if _, err := io.Copy(wc, src); err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error copying file to Firebase Storage: %v", err))
	}
	if err := wc.Close(); err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error closing writer: %v", err))
	}

	// Set ACL for public read access after creating the object
	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error setting ACL: %v", err))
	}

	// Get the download URL
	_, err = object.Attrs(context.Background())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file attributes: %v", err))
	}

	// Return the read-only URL to the client
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucketName, url.QueryEscape(storagePath))

	productRequest := new(web.ProductCreateRequest)

	if err := ctx.Bind(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	productTypeIDStr := ctx.FormValue("productTypeID")
	productTypeInt, err := strconv.Atoi(productTypeIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input product type"))
	}
	productTypeID := uint(productTypeInt)

	strAdminId := ctx.FormValue("adminID")
	adminIdInt, err := strconv.Atoi(strAdminId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input admin id"))
	}
	adminId := uint(adminIdInt)

	name := ctx.FormValue("name")

	ingredient := ctx.FormValue("ingredient")

	priceStr := ctx.FormValue("price")

	// Mengonversi string ke float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid price format"})
	}

	size := ctx.FormValue("size")

	productRequest.ProductTypeID = productTypeID
	productRequest.AdminID = adminId
	productRequest.Name = name
	productRequest.Ingredients = ingredient
	productRequest.Price = price
	productRequest.Size = size
	productRequest.Image = url

	result, err := c.ProductService.CreateProductService(ctx, *productRequest)

	result.ProductTypeID = productTypeID

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to create product"))
		}
	}

	response := res.ProductDomainToProductCreateResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success create product", response))

}

func (c *ProductHandlerImpl) UpdateProductHandler(ctx echo.Context) error {
	// Mendapatkan ID produk dari path parameter
	productIDStr := ctx.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid product ID"))
	}

	// Dapatkan data produk yang sudah ada
	existingProduct, err := c.ProductService.FindByIdProductService(ctx, uint(productID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to get existing product"))
	}

	// Memproses unggah gambar
	file, err := ctx.FormFile("image")
	var imageURL string

	if err != nil {
		// Jika tidak ada file gambar yang diunggah, gunakan gambar yang sudah ada (lama)
		imageURL = existingProduct.Image
	} else {
		// Jika ada file gambar yang diunggah, proses unggah gambar baru
		src, err := file.Open()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		defer src.Close()

		// Set the path to your service account JSON file
		serviceAccountKeyPath := "credentials.json"

		// Initialize Firebase Admin SDK
		opt := option.WithCredentialsFile(serviceAccountKeyPath)
		_, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error initializing app: %v", err))
		}

		// Set the destination path in Firebase Storage
		storagePath := "product/" + time.Now().Format("2006-01-02_15:04:05") + filepath.Ext(file.Filename)

		// Initialize Google Cloud Storage client
		client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(serviceAccountKeyPath))
		if err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error creating storage client: %v", err))
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
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unsupported file format: %s", fileExtension))
		}

		// Upload the file to Firebase Storage with the determined content type
		object := client.Bucket(bucketName).Object(storagePath)
		wc := object.NewWriter(context.Background())
		wc.ContentType = contentType

		if _, err := io.Copy(wc, src); err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error copying file to Firebase Storage: %v", err))
		}
		if err := wc.Close(); err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error closing writer: %v", err))
		}

		// Set ACL for public read access after creating the object
		if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error setting ACL: %v", err))
		}

		// Get the download URL
		_, err = object.Attrs(context.Background())
		if err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file attributes: %v", err))
		}

		// Return the read-only URL to the client
		imageURL = fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucketName, url.QueryEscape(storagePath))
	}

	// Mengonversi nilai-nilai dari request
	productTypeIDStr := ctx.FormValue("productTypeID")
	productTypeInt, _ := strconv.Atoi(productTypeIDStr)
	productTypeID := uint(productTypeInt)

	name := ctx.FormValue("name")
	ingredient := ctx.FormValue("ingredient")

	priceStr := ctx.FormValue("price")
	// Mengonversi string ke float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid price format"})
	}

	size := ctx.FormValue("size")

	// Mengupdate nilai-nilai produk yang sudah ada
	existingProduct.ProductTypeID = productTypeID
	existingProduct.Name = name
	existingProduct.Ingredients = ingredient
	existingProduct.Price = price
	existingProduct.Size = size
	existingProduct.Image = imageURL // Gunakan imageURL yang baru diunggah

	// Lakukan pembaruan data produk ke dalam database
	req := request.ProductDomainToProductUpdateRequest(existingProduct)
	result, err := c.ProductService.UpdateProductService(ctx, req, uint(productID))

	result.ID = existingProduct.ID

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to update product"))
		}
	}

	// Mengonversi hasil pembaruan ke dalam respons JSON
	response := res.ProductDomainToProductUpdateResponse(result)

	// Mengirimkan respons JSON sebagai tanggapan dari fungsi
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success update product", response))
}

func (c *ProductHandlerImpl) GetProductHandler(ctx echo.Context) error {

	productID := ctx.Param("id")
	productIDInt, err := strconv.Atoi(productID)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}

	product, err := c.ProductService.FindByIdProductService(ctx, uint(productIDInt))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMessage := "Get product data error"

		if strings.Contains(err.Error(), "product not found") {
			statusCode = http.StatusNotFound
			errorMessage = "Product not found"
		}

		return ctx.JSON(statusCode, helpers.ErrorResponse(errorMessage))
	}

	response := res.ProductDomainToProductResponse(product)

	responseCustom := res.ProductResponseToProductCostumResponse(response)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get product", responseCustom))
}

func (c *ProductHandlerImpl) GetProductsHandler(ctx echo.Context) error {
	result, err := c.ProductService.FindAllProductService(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product data error"))
	}

	response := res.ConvertProductResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all data product", response))
}

func (c *ProductHandlerImpl) GetProductByNameHandler(ctx echo.Context) error {
	productName := ctx.Param("name")

	result, err := c.ProductService.FindByNameProductService(ctx, productName)

	if err != nil {
		if strings.Contains(err.Error(), "failed to find products with the name") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to get product data by name"))
	}

	response := res.ConvertProductResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get product type by name", response))
}

func (c *ProductHandlerImpl) DeleteProductHandler(ctx echo.Context) error {
	productId := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	productIdUint := uint(productIdInt)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.ProductService.DeleteProductService(ctx, productIdUint)

	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data product error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data product", nil))
}
