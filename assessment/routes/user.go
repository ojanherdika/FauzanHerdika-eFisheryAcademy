package routes

import (
	"e-commerce/config"
	"e-commerce/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUser)
	e.POST("/login", handler.Login)

	r := e.Group("")
	config := middleware.JWTConfig{
		SigningKey: []byte(config.Config("SECRET")),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user", userHandler.GetUserByID)
	r.PUT("/user", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

}
