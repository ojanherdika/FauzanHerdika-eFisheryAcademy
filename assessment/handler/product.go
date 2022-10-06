package handler

import (
	"e-commerce/entity"
	"e-commerce/usecase"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CreateProduct struct {
	Name        string `form:"name"`
	Photo       string `form:"photo"`
	Price       int    `form:"price"`
	Category    string `form:"category"`
	Description string `form:"description"`
}
type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}
type ProductPrice struct {
	PriceMin int `json:"priceMin"`
	PriceMax int `json:"priceMax"`
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

// CreateProduct godoc
// @Summary Create a Product.
// @Description registering a user from full access.
// @Tags Product
// @Param Body formData CreateProduct true "Body Request fo Create Product"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products [post]
func (handler ProductHandler) CreateProduct(c echo.Context) error {
	// req := entity.CreateProductRequest{}
	// if err := c.Bind(&req); err != nil {
	// 	return err
	// }
	// product, err := handler.productUsecase.CreateProduct(req)
	// if err != nil {
	// 	return err
	// }
	// e := echo.New()
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "upload/product/", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	// filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())
	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())

	req := entity.CreateProductRequest{
		Name:        name,
		Photo:       filePath,
		Price:       price,
		Category:    category,
		Description: description,
	}
	product, err := handler.productUsecase.CreateProduct(req)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Create Product",
		"data":    product,
	})
}

// Get All Product godoc
// @Summary Get All Product.
// @Description get all Product.
// @Tags Product
// @Produce json
// @Success 200 {object} entity.ProductResponse
// @Router /products [get]
func (handler ProductHandler) GetAllProduct(c echo.Context) error {
	products, err := handler.productUsecase.GetAllProduct()
	//response handling with struct
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Product Failed",
			Error:   err.Error(),
		})
	}
	//response handling with new map interface
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Product",
		"data":    products,
	})
}

// Get Product by ID godoc
// @Summary Get Product by ID
// @Description Get Product by ID.
// @Tags Product
// @Produce json
// @Param id path int true "product id"
// @Success 200 {object} string
// @Router /product/{id} [get]
func (handler ProductHandler) GetProductByID(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))

	product, err := handler.productUsecase.GetProductById(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Product by ID Failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Product by ID",
		"data":    product,
	})
}

// Get Products by category godoc
// @Summary Get Products by category
// @Description Get Products by category.
// @Tags Product
// @Produce json
// @Param category path string true "product category"
// @Success 200 {object} string
// @Router /products/{category} [get]
func (handler ProductHandler) GetProductByCategory(c echo.Context) error {
	productCategory := c.Param("category")

	product, err := handler.productUsecase.GetProductByCategory(productCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Product by Category Failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Product by Category",
		"data":    product,
	})
}

// Get Products by price godoc
// @Summary Get Products by price
// @Description Get Products by price.
// @Tags Product
// @Produce json
// @Param Body priceMin,priceMax path string true "product price"
// @Success 200 {object} string
// @Router /products/{priceMin}/{priceMax} [get]
func (handler ProductHandler) GetProductByPrice(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.Param("priceMin"))
	priceMax, _ := strconv.Atoi(c.Param("priceMax"))

	product, err := handler.productUsecase.GetProductByPrice(priceMin, priceMax)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Product by Price Failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Product by Price",
		"data":    product,
	})
}

func (handler ProductHandler) UpdateProduct(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "upload/product/", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())

	req := entity.UpdateProductRequest{
		Name:        name,
		Photo:       filePath,
		Price:       price,
		Category:    category,
		Description: description,
	}
	productRequest := entity.UpdateProductRequest{}
	c.Bind(&productRequest)

	product, err := handler.productUsecase.UpdateProduct(req, productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update Product Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Update User",
		"data":    product,
	})
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete Product by id.
// @Tags Product
// @Produce json
// @Param id path int true "product id"
// @Success 200 {object} string
// @Router /products/{id} [delete]
func (handler ProductHandler) DeleteProduct(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	err := handler.productUsecase.DeleteProduct(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete Product Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "Delete Product Success")
}
