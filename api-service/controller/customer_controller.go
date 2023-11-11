package controller

import (
	"api-service/config"
	"api-service/helpers"
	"api-service/model"
	"api-service/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomerController interface {
	RequestGetProfile(ctx echo.Context) error
	RequestCreateLoan(ctx echo.Context) error
	RequestGetListLoan(ctx echo.Context) error
	RequestUpdateLoan(ctx echo.Context) error
	RequestCreateInstallment(ctx echo.Context) error
	RequestGetHistoryInstallment(ctx echo.Context) error
	RequestGetDetailLoan(ctx echo.Context) error
}

type customerControllerImpl struct {
	CustomerService service.CustomerService
	Validator       *validator.Validate
}

func NewCustomerController(newCustomerService service.CustomerService, newValidator *validator.Validate) CustomerController {
	return &customerControllerImpl{
		CustomerService: newCustomerService,
		Validator:       newValidator,
	}
}

func (controller *customerControllerImpl) RequestGetProfile(ctx echo.Context) error {

	message, statusCode, data := controller.CustomerService.ProcessGetProfile(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) RequestCreateLoan(ctx echo.Context) error {

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

	message, statusCode, data := controller.CustomerService.ProcessCreateLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) RequestGetListLoan(ctx echo.Context) error {

	message, statusCode, data := controller.CustomerService.ProcessGetListLoan(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) RequestUpdateLoan(ctx echo.Context) error {
	id := ctx.Param("id")

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

	message, statusCode, data := controller.CustomerService.ProcessUpdateLoan(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) RequestCreateInstallment(ctx echo.Context) error {
	var request model.RequestCreateInstallment
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

	message, statusCode, data := controller.CustomerService.ProcessCreateInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) RequestGetHistoryInstallment(ctx echo.Context) error {

	message, statusCode, data := controller.CustomerService.ProcessGetHistoryInstallment(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) RequestGetDetailLoan(ctx echo.Context) error {

	id := ctx.Param("id")

	message, statusCode, data := controller.CustomerService.ProcessGetDetailLoan(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
