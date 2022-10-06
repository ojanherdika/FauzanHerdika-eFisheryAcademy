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
	e.GET("/product/:id", productHandler.GetProductByID)
	e.GET("/products/:category", productHandler.GetProductByCategory)
	e.GET("/products/:priceMin/:priceMax", productHandler.GetProductByPrice)
}
