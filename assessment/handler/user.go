package handler

import (
	"e-commerce/entity"
	"e-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}
func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := handler.userUsecase.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Create user",
		"data":    user,
	})
}
func (handler UserHandler) GetAllUser(c echo.Context) error {
	users, err := handler.userUsecase.GetAllUser()
	//response handling with struct
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get User Failed",
			Error:   err.Error(),
		})
	}
	//response handling with new map interface
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All User",
		"data":    users,
	})
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(200, users)
}
func (handler UserHandler) GetUserByID(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := handler.userUsecase.GetUserById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get User by ID Failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get User by ID",
		"data":    user,
	})
}
func (handler UserHandler) UpdateUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	userRequest := entity.UpdateUserRequset{}
	c.Bind(&userRequest)

	user, err := handler.userUsecase.UpdateUser(userRequest, userId)
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
		"data":    user,
	})
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	err := handler.userUsecase.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete User Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "Delete User Berhasil")
}
