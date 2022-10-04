package handler

import (
	"e-commerce/entity"
	"e-commerce/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUsecase *usecase.CartUsecase
}

func NewCartHandler(cartUsecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase}
}
func (handler CartHandler) CreateCart(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"]

	strID := fmt.Sprintf("%v", userId)
	convertedUserID, _ := strconv.Atoi(strID)
	req := entity.CreateCartRequest{
		UserID:   convertedUserID,
		Checkout: false,
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	cart, err := handler.cartUsecase.CreateCart(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Create Cart",
		"data":    cart,
	})
}
func (handler CartHandler) GetAllCart(c echo.Context) error {
	carts, err := handler.cartUsecase.GetAllCart()
	//response handling with struct
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Cart Failed",
			Error:   err.Error(),
		})
	}
	//response handling with new map interface
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Cart",
		"data":    carts,
	})
}
func (handler CartHandler) GetCartByID(c echo.Context) error {
	cartId, _ := strconv.Atoi(c.Param("id"))

	cart, err := handler.cartUsecase.GetCartById(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Cart by ID Failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get User by ID",
		"data":    cart,
	})
}
func (handler CartHandler) UpdateCart(c echo.Context) error {
	cartId, _ := strconv.Atoi(c.Param("id"))
	cartRequest := entity.UpdateCartRequest{}
	c.Bind(&cartRequest)

	cart, err := handler.cartUsecase.UpdateCart(cartRequest, cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update Cart Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Update Cart",
		"data":    cart,
	})
}

func (handler CartHandler) DeleteCart(c echo.Context) error {
	cartId, _ := strconv.Atoi(c.Param("id"))
	err := handler.cartUsecase.DeleteCart(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete Cart Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "Delete Cart Berhasil")
}
