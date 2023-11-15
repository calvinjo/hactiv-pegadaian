package main

import (
	"api-service/config"
	"api-service/controller"
	"api-service/docs"
	"api-service/model"
	"api-service/repository"
	"api-service/service"

	_ "api-service/docs"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Setup Config
	env := config.New().SetArgs()
	validator := config.InitValidator()

	// Setup App
	UsersRepository := repository.NewUsersRepository(env.UsersServiceUrl)
	TransactionRepository := repository.NewTransactionRepository(env.TransactionsServiceUrl)

	AuthService := service.NewAuthService(UsersRepository)
	AuthController := controller.NewAuthController(AuthService, validator)

	AdminService := service.NewAdminService(UsersRepository, TransactionRepository)
	AdminController := controller.NewAdminController(AdminService, validator)

	CustomerService := service.NewCustomerService(UsersRepository, TransactionRepository)
	CustomerController := controller.NewCustomerController(CustomerService, validator)

	app := echo.New()
	app.Debug = true
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// @title Swagger API Gateway Service
	// @version 1.0
	// @description This is a sample server API Service Gateway.
	// @BasePath /api/v1
	docs.SwaggerInfo.Host = "0.0.0.0:" + env.AppsPort

	app.POST("/api/v1/register", func(ctx echo.Context) error {
		return AuthController.Register(ctx)
	})

	app.POST("/api/v1/login", func(ctx echo.Context) error {
		return AuthController.Login(ctx)
	})

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	//Setup JWT Echo
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtClaims)
		},
		SigningKey: []byte("mysecret"),
	}

	// Route Admin
	authAdmin := app.Group("/api/v1/admin")

	authAdmin.Use(echojwt.WithConfig(config))

	authAdmin.GET("/list-user", func(ctx echo.Context) error {
		return AdminController.RequestGetListUsers(ctx)
	}, restricted("admin"))

	authAdmin.GET("/detail-user/:id", func(ctx echo.Context) error {
		return AdminController.RequestGetDetailUsers(ctx)
	}, restricted("admin"))

	authAdmin.POST("/create-user", func(ctx echo.Context) error {
		return AdminController.RequestCreateUsers(ctx)
	}, restricted("admin"))

	authAdmin.PUT("/update-user/:id", func(ctx echo.Context) error {
		return AdminController.RequestUpdateUsers(ctx)
	}, restricted("admin"))

	authAdmin.DELETE("/delete-user/:id", func(ctx echo.Context) error {
		return AdminController.RequestDeleteUsers(ctx)
	}, restricted("admin"))

	authAdmin.GET("/list-loan", func(ctx echo.Context) error {
		return AdminController.RequestGetListLoan(ctx)
	}, restricted("admin"))

	authAdmin.GET("/detail-loan/:id", func(ctx echo.Context) error {
		return AdminController.RequestGetDetailLoan(ctx)
	}, restricted("admin"))

	authAdmin.PATCH("/update-status-loan/:id", func(ctx echo.Context) error {
		return AdminController.RequestUpdateStatusLoan(ctx)
	}, restricted("admin"))

	authAdmin.GET("/list-max-loan", func(ctx echo.Context) error {
		return AdminController.RequestGetListMaxLoan(ctx)
	}, restricted("admin"))

	authAdmin.GET("/detail-max-loan/:id", func(ctx echo.Context) error {
		return AdminController.RequestGetDetailMaxLoan(ctx)
	}, restricted("admin"))

	authAdmin.PUT("/update-max-loan/:id", func(ctx echo.Context) error {
		return AdminController.RequestUpdateMaxLoan(ctx)
	}, restricted("admin"))

	authAdmin.GET("/list-installment", func(ctx echo.Context) error {
		return AdminController.RequestGetListInstallment(ctx)
	}, restricted("admin"))

	//Route Customer
	authCustomer := app.Group("/api/v1/customer")

	authCustomer.Use(echojwt.WithConfig(config))

	authCustomer.GET("/profile", func(ctx echo.Context) error {
		return CustomerController.RequestGetProfile(ctx)
	}, restricted("customer"))

	authCustomer.POST("/create-loan", func(ctx echo.Context) error {
		return CustomerController.RequestCreateLoan(ctx)
	}, restricted("customer"))

	authCustomer.GET("/list-loan", func(ctx echo.Context) error {
		return CustomerController.RequestGetListLoan(ctx)
	}, restricted("customer"))

	authCustomer.PUT("/update-loan/:id", func(ctx echo.Context) error {
		return CustomerController.RequestUpdateLoan(ctx)
	}, restricted("customer"))

	authCustomer.POST("/create-installment", func(ctx echo.Context) error {
		return CustomerController.RequestCreateInstallment(ctx)
	}, restricted("customer"))

	authCustomer.GET("/history-installment", func(ctx echo.Context) error {
		return CustomerController.RequestGetHistoryInstallment(ctx)
	}, restricted("customer"))

	authCustomer.GET("/detail-loan/:id", func(ctx echo.Context) error {
		return CustomerController.RequestGetDetailLoan(ctx)
	}, restricted("customer"))

	app.Start(":" + env.AppsPort)

}

func restricted(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*model.JwtClaims)
			role := claims.Roles

			//set context
			c.Set("roles", role)
			c.Set("userId", claims.UserID)

			// Mengecek apakah pengguna memiliki peran yang diizinkan
			allowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					allowed = true
					break
				}
			}

			if !allowed {
				return echo.ErrForbidden
			}

			return next(c)
		}
	}
}
