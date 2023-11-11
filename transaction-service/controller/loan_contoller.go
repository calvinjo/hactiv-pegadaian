package controller

import (
	"log"
	"transaction-service/config"
	"transaction-service/helpers"
	"transaction-service/model"
	"transaction-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type LoanController interface {
	RequestAllLoan(ctx echo.Context) error
	RequestGetDetailLoan(ctx echo.Context) error
	RequestCreateLoan(ctx echo.Context) error
	RequestUpdateLoan(ctx echo.Context) error
	RequestUpdateStatusLoan(ctx echo.Context) error
	RequestDeleteLoan(ctx echo.Context) error
}

type loanControllerImpl struct {
	LoanService service.LoanService
	Validator   *validator.Validate
}

func NewLoanController(newLoanService service.LoanService, newValidator *validator.Validate) LoanController {
	return &loanControllerImpl{
		LoanService: newLoanService,
		Validator:   newValidator,
	}
}

func (controller *loanControllerImpl) RequestGetDetailLoan(ctx echo.Context) error {
	var request model.RequestGetDetailLoan
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	status, err := controller.validation(ctx, request)
	if !status {
		return err
	}

	message, statusCode, data := controller.LoanService.ProcessGetDetailLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *loanControllerImpl) RequestCreateLoan(ctx echo.Context) error {
	var request model.RequestCreateLoan
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	status, err := controller.validation(ctx, request)
	if !status {
		return err
	}

	message, statusCode, data := controller.LoanService.ProcessCreateLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *loanControllerImpl) RequestAllLoan(ctx echo.Context) error {
	var request model.RequestGetAllLoan
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	status, err := controller.validation(ctx, request)
	if !status {
		return err
	}

	message, statusCode, data := controller.LoanService.ProcessGetAllLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *loanControllerImpl) RequestUpdateLoan(ctx echo.Context) error {
	var request model.RequestUpdateLoan
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	status, err := controller.validation(ctx, request)
	if !status {
		return err
	}

	message, statusCode, data := controller.LoanService.ProcessUpdateLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *loanControllerImpl) RequestUpdateStatusLoan(ctx echo.Context) error {
	var request model.RequestUpdateStatusLoan
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	status, err := controller.validation(ctx, request)
	if !status {
		return err
	}

	message, statusCode, data := controller.LoanService.ProcessUpdateStatusLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *loanControllerImpl) RequestDeleteLoan(ctx echo.Context) error {
	var request model.RequestDeleteLoan
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	status, err := controller.validation(ctx, request)
	if !status {
		return err
	}

	message, statusCode, data := controller.LoanService.ProcessDeleteLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *loanControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
	err := controller.Validator.Struct(request)
	if err != nil {
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, dataField := range castedObject {
				fieldName := " " + dataField.Field()
				return false, helpers.GenerateResponse(ctx, config.InvalidFieldFormat, fieldName, nil, err)
			}
		}
	}

	return true, nil
}
