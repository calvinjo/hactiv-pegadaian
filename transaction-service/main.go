package main

import (
	"transaction-service/config"
	"transaction-service/controller"
	"transaction-service/repository"
	"transaction-service/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Setup Config
	env := config.New().SetArgs()
	validator := config.InitValidator()

	// Setup App
	PostgreSql := config.NewPostgreDatabase(env)

	LoanRepository := repository.NewLoanRepository(*PostgreSql)
	InstallmentRepository := repository.NewInstallmentRepository(*PostgreSql)
	MaxLoanRepository := repository.NewMaxLoanRepository(*PostgreSql)

	LoanService := service.NewLoanService(LoanRepository, MaxLoanRepository)
	LoanController := controller.NewLoanController(LoanService, validator)

	InstallmentService := service.NewInstallmentService(InstallmentRepository, LoanRepository, MaxLoanRepository)
	InstallmentController := controller.NewInstallmentController(InstallmentService, validator)

	MaxLoanService := service.NewMaxLoanService(MaxLoanRepository)
	MaxLoanController := controller.NewMaxLoanController(MaxLoanService, validator)

	app := echo.New()
	app.Debug = true

	//Route
	app.POST("/api/v1/detail-loan", func(ctx echo.Context) error {
		return LoanController.RequestGetDetailLoan(ctx)
	})

	app.POST("/api/v1/create-loan", func(ctx echo.Context) error {
		return LoanController.RequestCreateLoan(ctx)
	})

	app.POST("/api/v1/all-loan", func(ctx echo.Context) error {
		return LoanController.RequestAllLoan(ctx)
	})

	app.PUT("/api/v1/update-loan", func(ctx echo.Context) error {
		return LoanController.RequestUpdateLoan(ctx)
	})

	app.PATCH("/api/v1/update-status-loan", func(ctx echo.Context) error {
		return LoanController.RequestUpdateStatusLoan(ctx)
	})

	app.DELETE("/api/v1/delete-loan", func(ctx echo.Context) error {
		return LoanController.RequestDeleteLoan(ctx)
	})

	app.POST("/api/v1/detail-installment", func(ctx echo.Context) error {
		return InstallmentController.RequestGetDetailInstallment(ctx)
	})

	app.POST("/api/v1/create-installment", func(ctx echo.Context) error {
		return InstallmentController.RequestCreateInstallment(ctx)
	})

	app.POST("/api/v1/all-installment", func(ctx echo.Context) error {
		return InstallmentController.RequestAllInstallment(ctx)
	})

	app.PUT("/api/v1/update-installment", func(ctx echo.Context) error {
		return InstallmentController.RequestUpdateInstallment(ctx)
	})

	app.DELETE("/api/v1/delete-installment", func(ctx echo.Context) error {
		return InstallmentController.RequestDeleteInstallment(ctx)
	})

	app.POST("/api/v1/detail-max-loan", func(ctx echo.Context) error {
		return MaxLoanController.RequestGetDetailMaxLoan(ctx)
	})

	app.POST("/api/v1/create-max-loan", func(ctx echo.Context) error {
		return MaxLoanController.RequestCreateMaxLoan(ctx)
	})

	app.POST("/api/v1/all-max-loan", func(ctx echo.Context) error {
		return MaxLoanController.RequestGetAllMaxLoan(ctx)
	})

	app.PUT("/api/v1/update-max-loan", func(ctx echo.Context) error {
		return MaxLoanController.RequestUpdateMaxLoan(ctx)
	})

	app.DELETE("/api/v1/delete-max-loan", func(ctx echo.Context) error {
		return MaxLoanController.RequestDeleteMaxLoan(ctx)
	})

	app.Start(":" + env.AppsPort)
}
