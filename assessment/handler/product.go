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

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}
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

func (handler ProductHandler) DeleteProduct(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	err := handler.productUsecase.DeleteUser(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete Product Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "Delete Product Success")
}
