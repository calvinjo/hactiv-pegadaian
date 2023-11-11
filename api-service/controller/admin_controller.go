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

type AdminController interface {
	RequestGetListUsers(ctx echo.Context) error
	RequestGetDetailUsers(ctx echo.Context) error
	RequestCreateUsers(ctx echo.Context) error
	RequestUpdateUsers(ctx echo.Context) error
	RequestDeleteUsers(ctx echo.Context) error

	RequestGetListLoan(ctx echo.Context) error
	RequestGetDetailLoan(ctx echo.Context) error
	RequestUpdateStatusLoan(ctx echo.Context) error

	RequestGetListMaxLoan(ctx echo.Context) error
	RequestGetDetailMaxLoan(ctx echo.Context) error
	RequestUpdateMaxLoan(ctx echo.Context) error

	RequestGetListInstallment(ctx echo.Context) error
}

type adminControllerImpl struct {
	AdminService service.AdminService
	Validator    *validator.Validate
}

func NewAdminController(newAdminService service.AdminService, newValidator *validator.Validate) AdminController {
	return &adminControllerImpl{
		AdminService: newAdminService,
		Validator:    newValidator,
	}
}

func (controller *adminControllerImpl) RequestGetListUsers(ctx echo.Context) error {

	message, statusCode, data := controller.AdminService.ProcessGetListUsers(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestGetDetailUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessGetDetailUsers(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestCreateUsers(ctx echo.Context) error {
	var request model.RequestCreateUsers
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

	message, statusCode, data := controller.AdminService.ProcessCreateUsers(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestUpdateUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	var request model.RequestUpdateUsers
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

	message, statusCode, data := controller.AdminService.ProcessUpdateUsers(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestDeleteUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessDeleteUsers(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestGetListLoan(ctx echo.Context) error {
	var request model.RequestListLoan
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

	message, statusCode, data := controller.AdminService.ProcessGetListLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestGetDetailLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessGetDetailLoan(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestUpdateStatusLoan(ctx echo.Context) error {
	id := ctx.Param("id")

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

	message, statusCode, data := controller.AdminService.ProcessUpdateStatusLoan(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestGetListMaxLoan(ctx echo.Context) error {
	var request model.RequestListMaxLoan
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

	message, statusCode, data := controller.AdminService.ProcessGetListMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestGetDetailMaxLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessGetDetailMaxLoan(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestUpdateMaxLoan(ctx echo.Context) error {
	id := ctx.Param("id")

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

	message, statusCode, data := controller.AdminService.ProcessUpdateMaxLoan(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) RequestGetListInstallment(ctx echo.Context) error {

	var request model.RequestListInstallment
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

	message, statusCode, data := controller.AdminService.ProcessGetListInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
