package handler

import (
	"fmt"
	"net/http"
	"os"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	"qbills/utils/helpers/firebase"
	"qbills/utils/helpers/middleware"
	"qbills/utils/request"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProductHandler interface {
	CreateProductHandler(ctx echo.Context) error
	UpdateProductHandler(ctx echo.Context) error
	GetProductHandler(ctx echo.Context) error
	GetProductsHandler(ctx echo.Context) error
	GetProductByNameHandler(ctx echo.Context) error
	GetProductByCategoryHandler(ctx echo.Context) error
	DeleteProductHandler(ctx echo.Context) error
	FindPaginationProduct(ctx echo.Context) error
	ProductAIHandler(ctx echo.Context) error
	GetProductNames(ctx echo.Context) (map[uint]helpers.ProductDataAIRecommended, error)
	GetBestProductsHandler(ctx echo.Context) error
}

type ProductHandlerImpl struct {
	ProductService services.ProductService
}

func NewProductHandler(ProductService services.ProductService) ProductHandler {
	return &ProductHandlerImpl{
		ProductService: ProductService,
	}
}

func (c *ProductHandlerImpl) CreateProductHandler(ctx echo.Context) error {
	adminId := middleware.ExtractTokenAdminId(ctx)
	if adminId == 0.0 {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid token admin"))
	}
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
	}

	productTypeIDStr := ctx.FormValue("productTypeId")
	productTypeInt, err := strconv.Atoi(productTypeIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input product type"))
	}
	productTypeID := uint(productTypeInt)

	name := ctx.FormValue("name")
	ingredients := ctx.FormValue("ingredients")
	price := ctx.FormValue("productDetail[price]")
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid client input price"))
	}

	productRequest.ProductTypeID = productTypeID
	productRequest.AdminID = uint(adminId)
	productRequest.Name = name
	productRequest.Ingredients = ingredients
	productRequest.Image = url
	productRequest.ProductDetail.Price = priceFloat

	result, err := c.ProductService.CreateProductService(ctx, *productRequest)

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

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Success create product", response))
}

func (c *ProductHandlerImpl) UpdateProductHandler(ctx echo.Context) error {
	// Mendapatkan ID produk dari path parameter
	productIDStr := ctx.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid product ID"))
	}

	// Dapatkan data produk yang sudah ada
	existingProduct, err := c.ProductService.FindByIdProductService(ctx, uint(productID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to get existing product"))
	}

	// Cek apakah ada file yang diunggah oleh klien
	_, err = ctx.FormFile("image")
	var imageURL string
	if err != nil {
		// Jika tidak ada file yang diunggah, gunakan gambar yang sudah ada (lama)
		imageURL = existingProduct.Image
	} else {
		// Jika ada file yang diunggah, proses unggah gambar baru
		// Lakukan proses upload gambar baru ke dalam sistem (misalnya menggunakan firebase)
		uploadedURL, err := firebase.UploadImageProduct(ctx)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to upload file"))
		}
		imageURL = uploadedURL
	}

	// Mengonversi nilai-nilai dari request
	productTypeIDStr := ctx.FormValue("productTypeId")
	productTypeInt, _ := strconv.Atoi(productTypeIDStr)
	productTypeID := uint(productTypeInt)

	name := ctx.FormValue("name")
	ingredients := ctx.FormValue("ingredients")

	// Mengupdate nilai-nilai produk yang sudah ada
	existingProduct.ProductTypeID = productTypeID
	existingProduct.Name = name
	existingProduct.Ingredients = ingredients
	existingProduct.Image = imageURL // Gunakan imageURL yang baru diunggah

	// Lakukan pembaruan data produk ke dalam database
	req := request.ProductDomainToProductUpdateRequest(existingProduct)
	result, err := c.ProductService.UpdateProductService(ctx, req, uint(productID))
	result.ID = existingProduct.ID

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to update product"))
		}
	}

	// Mengonversi hasil pembaruan ke dalam respons JSON
	response := res.ProductDomainToProductUpdateResponse(result)

	// Mengirimkan respons JSON sebagai tanggapan dari fungsi
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success update product", response))
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

	response := res.ProductsDomainToProductsResponse(product)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get product", response))
}

func (c *ProductHandlerImpl) GetProductsHandler(ctx echo.Context) error {
	products, totalProducts, err := c.ProductService.FindAllProductService(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product data error"))
	}

	response := res.ConvertProductResponse(products)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponseWithTotal("success get all data product", response, totalProducts))
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

func (c *ProductHandlerImpl) GetProductByCategoryHandler(ctx echo.Context) error {
	productTypeID := ctx.Param("productTypeId")
	productTypeIDUint64, err := strconv.ParseUint(productTypeID, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to parse product type"))
	}
	result, err := c.ProductService.FindByCategoryProductService(ctx, uint(productTypeIDUint64))

	if err != nil {
		if strings.Contains(err.Error(), "failed to find products with the category") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to get product data by category"))
	}

	response := res.ConvertProductResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get product by category", response))
}

func (c *ProductHandlerImpl) DeleteProductHandler(ctx echo.Context) error {
	productId := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	productIdUint := uint(productIdInt)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.ProductService.DeleteProductService(ctx, productIdUint)
	fmt.Println(err)
	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data product error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data product", nil))
}

func (c *ProductHandlerImpl) FindPaginationProduct(ctx echo.Context) error {

	response, meta, err := c.ProductService.FindPaginationProduct(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "Product is empty") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed get pagination product"))
	}

	productResponse := res.ConvertProductResponse(response)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponseWithMeta("succesfully get data product", productResponse, meta))
}

func (c *ProductHandlerImpl) GetProductNames(ctx echo.Context) (map[uint]helpers.ProductDataAIRecommended, error) {
	productMap := make(map[uint]helpers.ProductDataAIRecommended)

	products, _, err := c.ProductService.FindAllProductService(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return productMap, ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product not found"))
		}

		return productMap, ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product data error"))
	}

	for _, product := range products {
		productData := helpers.ProductDataAIRecommended{
			Name:        product.Name,
			Ingredients: product.Ingredients,
		}
		productMap[uint(product.ID)] = productData
	}

	return productMap, nil
}

func (c *ProductHandlerImpl) ProductAIHandler(ctx echo.Context) error {
	userInput := web.ProductRecommendRequest{}
	err := ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}
	openAIKey := os.Getenv("OPEN_AI_KEY")
	productMap, err := c.GetProductNames(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Error getting product names"))
	}

	result, err := helpers.ProductAI(productMap, openAIKey, userInput.Input)
	if err != nil {
		log.Error("Error calling ProductAI:", err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Error getting product recommendation"))
	}

	log.Info("ProductAI Result:", result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get product recommendation", result))
}


func (c *ProductHandlerImpl) GetBestProductsHandler(ctx echo.Context) error {
	bestProducts, err := c.ProductService.FindBestSellingProduct()

	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("best product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get best product data error"))
	}

	response := res.ConvertBestProductResponse(bestProducts)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all data best product", response))
}
