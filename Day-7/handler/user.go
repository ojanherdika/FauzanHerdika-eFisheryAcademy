package handler

import (
	"clean_arch/entity"
	"clean_arch/usecase"
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
	req := entity.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := handler.userUsecase.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(201, user)
}
func (handler UserHandler) GetAllUser(c echo.Context) error {
	users, err := handler.userUsecase.GetAllUser()

	if err != nil {
		return err
	}
	return c.JSON(200, users)
}
func (handler UserHandler) GetUserByID(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := handler.userUsecase.GetUserById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Get user by id failed")
	}

	return c.JSON(200, user)
}
