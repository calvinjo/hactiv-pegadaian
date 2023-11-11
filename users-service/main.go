package main

import (
	"users-service/config"
	"users-service/controller"
	"users-service/repository"
	"users-service/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Setup Config
	env := config.New().SetArgs()
	validator := config.InitValidator()

	// Setup App
	PostgreSql := config.NewPostgreDatabase(env)
	UsersRepository := repository.NewUsersRepository(*PostgreSql)
	UsersService := service.NewUsersService(UsersRepository)
	UsersController := controller.NewUsersController(UsersService, validator)

	app := echo.New()
	app.Debug = true

	//Route
	app.POST("/api/v1/detail-user", func(ctx echo.Context) error {
		return UsersController.RequestGetDetailUser(ctx)
	})

	app.POST("/api/v1/create-user", func(ctx echo.Context) error {
		return UsersController.RequestCreateUser(ctx)
	})

	app.GET("/api/v1/all-user", func(ctx echo.Context) error {
		return UsersController.RequestAllUser(ctx)
	})

	app.PUT("/api/v1/update-user", func(ctx echo.Context) error {
		return UsersController.RequestUpdateUser(ctx)
	})

	app.DELETE("/api/v1/delete-user", func(ctx echo.Context) error {
		return UsersController.RequestDeleteUser(ctx)
	})

	app.Start(":" + env.AppsPort)
}
