package handler

import (
<<<<<<< Updated upstream
	"github.com/labstack/echo/v4"
=======
<<<<<<< Updated upstream
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"io"
=======
<<<<<<< Updated upstream
	"github.com/labstack/echo/v4"
=======
<<<<<<< Updated upstream
	"github.com/labstack/echo/v4"
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
<<<<<<< Updated upstream
	"qbills/utils/helpers/firebase"
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	"qbills/utils/helpers/firebase"
=======
<<<<<<< Updated upstream
=======
	"qbills/utils/helpers/firebase"
<<<<<<< Updated upstream
=======
	"qbills/utils/helpers/middleware"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	"qbills/utils/request"
	res "qbills/utils/response"
	"strconv"
	"strings"
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	"time"
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
)

type ProductHandler interface {
	CreateProductHandler(ctx echo.Context) error
	UpdateProductHandler(ctx echo.Context) error
	GetProductHandler(ctx echo.Context) error
	GetProductsHandler(ctx echo.Context) error
	GetProductByNameHandler(ctx echo.Context) error
<<<<<<< Updated upstream
	GetProductByCategoryHandler(ctx echo.Context) error
	DeleteProductHandler(ctx echo.Context) error
	FindPaginationProduct(ctx echo.Context) error
=======
<<<<<<< Updated upstream
	DeleteProductHandler(ctx echo.Context) error
=======
	GetProductByCategoryHandler(ctx echo.Context) error
	DeleteProductHandler(ctx echo.Context) error
	FindPaginationProduct(ctx echo.Context) error
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}

type ProductHandlerImpl struct {
	ProductService services.ProductService
}

func NewProductHandler(ProductService services.ProductService) ProductHandler {
	return &ProductHandlerImpl{ProductService: ProductService}
}

func (c *ProductHandlerImpl) CreateProductHandler(ctx echo.Context) error {
<<<<<<< Updated upstream
	url, err := firebase.UploadImageProduct(ctx)

=======
<<<<<<< Updated upstream

<<<<<<< Updated upstream
	// Set the path to your service account JSON file
	serviceAccountKeyPath := "credentials.json"
=======
<<<<<<< Updated upstream
	url, err := firebase.UploadImageProduct(ctx)
=======
	url, err := c.ProductService.UploadImageProduct(ctx)
=======
<<<<<<< Updated upstream

<<<<<<< Updated upstream
	url, err := c.ProductService.UploadImageProduct(ctx)
=======
	url, err := firebase.UploadImageProduct(ctx)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	_, err := firebase.NewApp(context.Background(), nil, opt)
>>>>>>> Stashed changes
	if err != nil {
		if url == "" {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("File not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to upload file"))
	}
<<<<<<< Updated upstream

<<<<<<< Updated upstream
	productRequest := new(web.ProductCreateRequest)

	if err := ctx.Bind(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input"))
=======
<<<<<<< Updated upstream
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

=======
=======
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	productRequest := new(web.ProductCreateRequest)

	if err := ctx.Bind(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
<<<<<<< Updated upstream
=======
=======
	admin := middleware.ExtractTokenAdminId(ctx)
	url, err := firebase.UploadImageProduct(ctx)

	if err != nil {
		if url == "" {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("File not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to upload file"))
	}
	
	productRequest := new(web.ProductCreateRequest)

	if err := ctx.Bind(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input"))
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	}

	productTypeIDStr := ctx.FormValue("productTypeID")
	productTypeInt, err := strconv.Atoi(productTypeIDStr)
	if err != nil {
<<<<<<< Updated upstream
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input product type"))
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input product type"))
>>>>>>> Stashed changes
	}
	productTypeID := uint(productTypeInt)

	strAdminId := ctx.FormValue("adminID")
	adminIdInt, err := strconv.Atoi(strAdminId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input admin id"))
	}
	adminId := uint(adminIdInt)

	name := ctx.FormValue("name")
<<<<<<< Updated upstream
	ingredients := ctx.FormValue("ingredients")
=======

<<<<<<< Updated upstream
	description := ctx.FormValue("description")
=======
<<<<<<< Updated upstream
	ingredients := ctx.FormValue("ingredients")
=======
<<<<<<< Updated upstream
	ingredient := ctx.FormValue("ingredient")
=======
	ingredients := ctx.FormValue("ingredients")
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	priceStr := ctx.FormValue("price")

	// Mengonversi string ke float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid price format"})
	}

	stock := ctx.FormValue("stock")
	stockValue, _ := strconv.ParseUint(stock, 10, 64)
	stockUint := uint(stockValue)

	size := ctx.FormValue("size")
>>>>>>> Stashed changes

	productRequest.ProductTypeID = productTypeID
	productRequest.AdminID = adminId
	productRequest.Name = name
<<<<<<< Updated upstream
	productRequest.Ingredients = ingredients
=======
<<<<<<< Updated upstream
	productRequest.Description = description
=======
<<<<<<< Updated upstream
	productRequest.Ingredients = ingredients
=======
<<<<<<< Updated upstream
	productRequest.Ingredients = ingredient
>>>>>>> Stashed changes
	productRequest.Price = price
	productRequest.Size = size
=======
	productRequest.Ingredients = ingredients
>>>>>>> Stashed changes
	productRequest.Price = price
	productRequest.Stock = stockUint
	productRequest.Size = size
=======
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input product type"))
	}
	productTypeID := uint(productTypeInt)

	name := ctx.FormValue("name")
	ingredients := ctx.FormValue("ingredients")

	productRequest.ProductTypeID = productTypeID
	productRequest.AdminID = uint(admin)
	productRequest.Name = name
	productRequest.Ingredients = ingredients
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	productRequest.Image = url

	result, err := c.ProductService.CreateProductService(ctx, *productRequest)

<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
	result.ProductTypeID = productTypeID

>>>>>>> Stashed changes
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to create product"))
		}
	}

	result.ProductTypeID = productTypeID
	response := res.ProductDomainToProductCreateResponse(result)

<<<<<<< Updated upstream
	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Success create product", response))
=======
	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success create product", response))

<<<<<<< Updated upstream
=======
=======
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		case strings.Contains(err.Error(), "alpha"): 
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("name product is not valid must contain only alphabetical characters"))
		default:
			logrus.Error(err.Error())
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to create product"))
		}
	}

	result.ProductTypeID = productTypeID
	response := res.ProductDomainToProductCreateResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Success create product", response))
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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

	var imageURL string

	if err != nil {
		// Jika tidak ada file gambar yang diunggah, gunakan gambar yang sudah ada (lama)
		imageURL = existingProduct.Image
	} else {
		// Jika ada file gambar yang diunggah, proses unggah gambar baru
<<<<<<< Updated upstream
		url, err := firebase.UploadImageProduct(ctx)
=======
<<<<<<< Updated upstream
		src, err := file.Open()
=======
<<<<<<< Updated upstream
		url, err := c.ProductService.UploadImageProduct(ctx)
=======
<<<<<<< Updated upstream
		url, err := firebase.UploadImageProduct(ctx)
=======
<<<<<<< Updated upstream
		url, err := c.ProductService.UploadImageProduct(ctx)
=======
		url, err := firebase.UploadImageProduct(ctx)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed Upload file"))
		}

		imageURL = url

	}

	// Mengonversi nilai-nilai dari request
	productTypeIDStr := ctx.FormValue("productTypeID")
	productTypeInt, _ := strconv.Atoi(productTypeIDStr)
	productTypeID := uint(productTypeInt)

	name := ctx.FormValue("name")
<<<<<<< Updated upstream

	ingredients := ctx.FormValue("ingredients")
=======
<<<<<<< Updated upstream
	description := ctx.FormValue("description")
=======
<<<<<<< Updated upstream

	ingredients := ctx.FormValue("ingredients")
=======
<<<<<<< Updated upstream
	ingredient := ctx.FormValue("ingredient")
>>>>>>> Stashed changes

=======

	ingredients := ctx.FormValue("ingredients")
>>>>>>> Stashed changes

<<<<<<< Updated upstream
>>>>>>> Stashed changes
	priceStr := ctx.FormValue("price")
	// Mengonversi string ke float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid price format"})
	}

	stock := ctx.FormValue("stock")
	stockValue, _ := strconv.ParseUint(stock, 10, 64)
	stockUint := uint(stockValue)

	size := ctx.FormValue("size")
>>>>>>> Stashed changes

<<<<<<< Updated upstream
	// Mengupdate nilai-nilai produk yang sudah ada
	existingProduct.ProductTypeID = productTypeID
	existingProduct.Name = name
<<<<<<< Updated upstream
	existingProduct.Ingredients = ingredients
=======
	existingProduct.Description = description
	existingProduct.Price = price
	existingProduct.Stock = stockUint
	existingProduct.Size = size
=======
=======
>>>>>>> Stashed changes
	// Mengupdate nilai-nilai produk yang sudah ada
	existingProduct.ProductTypeID = productTypeID
	existingProduct.Name = name
	existingProduct.Ingredients = ingredients
<<<<<<< Updated upstream
	existingProduct.Price = price
	existingProduct.Size = size
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	existingProduct.Image = imageURL // Gunakan imageURL yang baru diunggah

	// Lakukan pembaruan data produk ke dalam database
	req := request.ProductDomainToProductUpdateRequest(existingProduct)
<<<<<<< Updated upstream

=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream

=======
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	result, err := c.ProductService.UpdateProductService(ctx, req, uint(productID))

	result.ID = existingProduct.ID

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		default:
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
			logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

=======
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

=======
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to get product data by name"))
	}

	response := res.ConvertProductResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get product type by name", response))
}

<<<<<<< Updated upstream
func (c *ProductHandlerImpl) GetProductByCategoryHandler(ctx echo.Context) error {
	productTypeID := ctx.Param("productTypeID")
	productTypeIDUint64, err := strconv.ParseUint(productTypeID, 10, 64)

	result, err := c.ProductService.FindByCategoryProductService(ctx, uint(productTypeIDUint64))

=======
<<<<<<< Updated upstream
=======
func (c *ProductHandlerImpl) GetProductByCategoryHandler(ctx echo.Context) error {
	productTypeID := ctx.Param("productTypeID")
	productTypeIDUint64, err := strconv.ParseUint(productTypeID, 10, 64)
<<<<<<< Updated upstream

	result, err := c.ProductService.FindByCategoryProductService(ctx, uint(productTypeIDUint64))

=======
	if err != nil {
		return err
	}

	result, err := c.ProductService.FindByCategoryProductService(ctx, uint(productTypeIDUint64))
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	if err != nil {
		if strings.Contains(err.Error(), "failed to find products with the category") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to get product data by category"))
	}

	response := res.ConvertProductResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get product by category", response))
}

<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

=======
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data product error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data product", nil))
}
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes

func (c *ProductHandlerImpl) FindPaginationProduct(ctx echo.Context) error {

	response, meta, err := c.ProductService.FindPaginationProduct(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "Product is empty") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}
<<<<<<< Updated upstream

=======
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed get pagination product"))
	}

	productResponse := res.ConvertProductResponse(response)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponseWithMeta("succesfully get data product", productResponse, meta))
}
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
