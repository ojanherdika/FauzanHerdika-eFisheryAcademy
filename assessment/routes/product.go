package routes

import (
	"e-commerce/handler"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, productHandler *handler.ProductHandler) {
	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetAllProduct)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)
	e.GET("/products/:id", productHandler.GetProductByID)
	// r := e.Group("/restricted")
	// config := middleware.JWTConfig{
	// 	SigningKey: []byte("SECRET"),
	// }
	// r.Use(middleware.JWTWithConfig(config))

	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {})
	// middware := e
	// middware.Use(midd.IsLogIn(e))

	// e.Use(middleware.JWTWithConfig())
}
