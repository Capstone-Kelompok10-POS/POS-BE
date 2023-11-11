package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
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
	svc, err := helpers.ConnectAWS()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	defer src.Close()

	uniqueFilename := uuid.New().String() + "_" + time.Now().Format("20060102150405") + filepath.Ext(file.Filename)

	fileExtension := filepath.Ext(uniqueFilename)
	contentType := "application/octet-stream" // Nilai default jika ekstensi tidak dikenali

	// Daftar ekstensi gambar yang dikenali
	imageExtensions := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		// Tambahkan ekstensi lain jika diperlukan
	}

	if val, ok := imageExtensions[fileExtension]; ok {
		contentType = val
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(uniqueFilename),
		Body:   src,
		// Gunakan tipe konten yang ditentukan oleh sistem (berdasarkan ekstensi)
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"), // Set ACL ke public-read
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// Dapatkan URL file yang diunggah
	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("AWS_BUCKET_NAME"), uniqueFilename)
	productRequest := new(web.ProductCreateRequest)

	if err := ctx.Bind(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	productTypeIDStr := ctx.FormValue("productTypeID")
	productTypeInt, _ := strconv.Atoi(productTypeIDStr)
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

	description := ctx.FormValue("description")

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

	productRequest.ProductTypeID = productTypeID
	productRequest.AdminID = adminId
	productRequest.Name = name
	productRequest.Description = description
	productRequest.Price = price
	productRequest.Stock = stockUint
	productRequest.Size = size
	productRequest.Image = imageURL

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
	svc, err := helpers.ConnectAWS()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

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

		uniqueFilename := uuid.New().String() + "_" + time.Now().Format("20060102150405") + filepath.Ext(file.Filename)

		fileExtension := filepath.Ext(uniqueFilename)
		contentType := "application/octet-stream" // Nilai default jika ekstensi tidak dikenali

		// Daftar ekstensi gambar yang dikenali
		imageExtensions := map[string]string{
			".jpg":  "image/jpeg",
			".jpeg": "image/jpeg",
			".png":  "image/png",
			// Tambahkan ekstensi lain jika diperlukan
		}

		if val, ok := imageExtensions[fileExtension]; ok {
			contentType = val
		}

		params := &s3.PutObjectInput{
			Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
			Key:         aws.String(uniqueFilename),
			Body:        src,
			ContentType: aws.String(contentType),
			ACL:         aws.String("public-read"),
		}

		_, err = svc.PutObject(params)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		// Dapatkan URL file yang diunggah
		imageURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("AWS_BUCKET_NAME"), uniqueFilename)
	}

	productTypeIDStr := ctx.FormValue("productTypeID")
	productTypeInt, _ := strconv.Atoi(productTypeIDStr)

	productTypeID := uint(productTypeInt)

	name := ctx.FormValue("name")

	description := ctx.FormValue("description")

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

	// Mendapatkan nilai-nilai dari request
	productRequest := &web.ProductUpdateRequest{
		ProductTypeID: productTypeID,
		Name:          name,
		Description:   description,
		Price:         price,
		Stock:         stockUint,
		Size:          size,
		Image:         imageURL,
	}

	// Mengupdate nilai-nilai produk yang sudah ada
	existingProduct.ProductTypeID = productRequest.ProductTypeID
	existingProduct.Name = productRequest.Name
	existingProduct.Description = productRequest.Description
	existingProduct.Price = productRequest.Price
	existingProduct.Stock = productRequest.Stock
	existingProduct.Size = productRequest.Size
	existingProduct.Image = productRequest.Image

	req := request.ProductDomainToProductUpdateRequest(existingProduct)

	// Lakukan pembaruan data produk ke dalam database
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

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get product", response))
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
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product data by name error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succsess get product type by name", result))
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
