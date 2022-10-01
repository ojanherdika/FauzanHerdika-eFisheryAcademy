package routes

import (
	"e-commerce/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUser)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.POST("/login", handler.Login)
}
