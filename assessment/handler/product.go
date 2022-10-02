package handler

import (
	"e-commerce/entity"
	"e-commerce/usecase"
	"net/http"
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
	req := entity.CreateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	product, err := handler.productUsecase.CreateProduct(req)
	if err != nil {
		return err
	}
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
	productRequest := entity.UpdateProductRequset{}
	c.Bind(&productRequest)

	product, err := handler.productUsecase.UpdateProduct(productRequest, productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update User Failed",
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
	return c.JSON(http.StatusOK, "Delete Product Berhasil")
}
