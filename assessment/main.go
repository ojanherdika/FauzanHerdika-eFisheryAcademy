package main

import (
	"e-commerce/config"
	"e-commerce/docs"
	"e-commerce/handler"
	"e-commerce/repository"
	"e-commerce/routes"
	"e-commerce/usecase"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "E-Commerce"
	docs.SwaggerInfo.Description = "Assessment Associate Software Engineer eFishery Academy Batch 2.0"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	config.Database()
	config.Migrate()
	e := echo.New()
	e.Static("/upload/product", "upload/product/")
	e.Static("/upload/payment", "upload/payment/")
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := usecase.NewCartUseCase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUsecase)
	routes.Routes(e, userHandler)
	routes.ProductRoutes(e, productHandler)
	routes.CartRoutes(e, cartHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
