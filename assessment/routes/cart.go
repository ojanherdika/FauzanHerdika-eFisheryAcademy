package routes

import (
	"e-commerce/config"
	"e-commerce/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CartRoutes(e *echo.Echo, cartHandler *handler.CartHandler) {
	e.GET("/carts", cartHandler.GetAllCart)
	r := e.Group("")
	config := middleware.JWTConfig{
		SigningKey: []byte(config.Config("SECRET")),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/carts/:id", cartHandler.GetCartByID)
	r.GET("/carts/user", cartHandler.GetCartByUser)
	r.POST("/carts", cartHandler.CreateCart)
	r.PUT("/carts/:id", cartHandler.UpdateCart)
	r.DELETE("/carts/:id", cartHandler.DeleteCart)
	r.POST("/payment", cartHandler.Payment)

}
