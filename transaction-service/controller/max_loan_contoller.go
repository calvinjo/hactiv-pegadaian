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

type MaxLoanController interface {
	RequestGetAllMaxLoan(ctx echo.Context) error
	RequestGetDetailMaxLoan(ctx echo.Context) error
	RequestCreateMaxLoan(ctx echo.Context) error
	RequestUpdateMaxLoan(ctx echo.Context) error
	RequestDeleteMaxLoan(ctx echo.Context) error
}

type maxLoanControllerImpl struct {
	MaxLoanService service.MaxLoanService
	Validator      *validator.Validate
}

func NewMaxLoanController(newMaxLoanService service.MaxLoanService, newValidator *validator.Validate) MaxLoanController {
	return &maxLoanControllerImpl{
		MaxLoanService: newMaxLoanService,
		Validator:      newValidator,
	}
}

func (controller *maxLoanControllerImpl) RequestGetDetailMaxLoan(ctx echo.Context) error {
	var request model.RequestGetDetailMaxLoan
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

	message, statusCode, data := controller.MaxLoanService.ProcessGetDetailMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *maxLoanControllerImpl) RequestCreateMaxLoan(ctx echo.Context) error {
	var request model.RequestCreateMaxLoan
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

	message, statusCode, data := controller.MaxLoanService.ProcessCreateMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *maxLoanControllerImpl) RequestGetAllMaxLoan(ctx echo.Context) error {
	var request model.RequestGetAllMaxLoan
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

	//Validation
	message, statusCode, data := controller.MaxLoanService.ProcessGetAllMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *maxLoanControllerImpl) RequestUpdateMaxLoan(ctx echo.Context) error {
	var request model.RequestUpdateMaxLoan
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

	message, statusCode, data := controller.MaxLoanService.ProcessUpdateMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *maxLoanControllerImpl) RequestDeleteMaxLoan(ctx echo.Context) error {
	var request model.RequestDeleteMaxLoan
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

	message, statusCode, data := controller.MaxLoanService.ProcessDeleteMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *maxLoanControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
