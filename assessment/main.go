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
	e := echo.New()
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	routes.Routes(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
	// config.Migrate()
}
