package routes

import (
	"github.com/labstack/echo/v4"

	"clean_arch/handler"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/createUser", userHandler.CreateUser)
	e.GET("/getUsers", userHandler.GetAllUser)
}
