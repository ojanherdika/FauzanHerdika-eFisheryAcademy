package main

import (
	"e-commerce/config"
	"e-commerce/handler"
	"e-commerce/repository"
	"e-commerce/routes"
	"e-commerce/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	// config.Migrate()
	e := echo.New()
	e.Static("/upload/product", "upload/product/")
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	routes.Routes(e, userHandler)
	routes.ProductRoutes(e, productHandler)
	// routes.StaticRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))

}
