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

type InstallmentController interface {
	RequestAllInstallment(ctx echo.Context) error
	RequestGetDetailInstallment(ctx echo.Context) error
	RequestCreateInstallment(ctx echo.Context) error
	RequestUpdateInstallment(ctx echo.Context) error
	RequestDeleteInstallment(ctx echo.Context) error
}

type installmentControllerImpl struct {
	InstallmentService service.InstallmentService
	Validator          *validator.Validate
}

func NewInstallmentController(newInstallmentService service.InstallmentService, newValidator *validator.Validate) InstallmentController {
	return &installmentControllerImpl{
		InstallmentService: newInstallmentService,
		Validator:          newValidator,
	}
}

func (controller *installmentControllerImpl) RequestGetDetailInstallment(ctx echo.Context) error {
	var request model.RequestGetDetailInstallment
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

	message, statusCode, data := controller.InstallmentService.ProcessGetDetailInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *installmentControllerImpl) RequestCreateInstallment(ctx echo.Context) error {
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

	message, statusCode, data := controller.InstallmentService.ProcessCreateInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *installmentControllerImpl) RequestAllInstallment(ctx echo.Context) error {
	var request model.RequestGetAllInstallment
	err := ctx.Bind(&request)
	if err != nil {
		log.Println(err)
		return helpers.GenerateResponse(ctx, config.InternalServerError, "", nil, err)
	}

	//Validation
	message, statusCode, data := controller.InstallmentService.ProcessGetAllInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *installmentControllerImpl) RequestUpdateInstallment(ctx echo.Context) error {
	var request model.RequestUpdateInstallment
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

	message, statusCode, data := controller.InstallmentService.ProcessUpdateInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *installmentControllerImpl) RequestDeleteInstallment(ctx echo.Context) error {
	var request model.RequestDeleteInstallment
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

	message, statusCode, data := controller.InstallmentService.ProcessDeleteInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *installmentControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
