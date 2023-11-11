package controller

import (
	"log"
	"users-service/config"
	"users-service/helpers"
	"users-service/model"
	"users-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UsersController interface {
	RequestAllUser(ctx echo.Context) error
	RequestGetDetailUser(ctx echo.Context) error
	RequestCreateUser(ctx echo.Context) error
	RequestUpdateUser(ctx echo.Context) error
	RequestDeleteUser(ctx echo.Context) error
}

type usersControllerImpl struct {
	UsersService service.UsersService
	Validator    *validator.Validate
}

func NewUsersController(newUsersService service.UsersService, newValidator *validator.Validate) UsersController {
	return &usersControllerImpl{
		UsersService: newUsersService,
		Validator:    newValidator,
	}
}

func (controller *usersControllerImpl) RequestGetDetailUser(ctx echo.Context) error {
	var request model.RequestGetDetailUser
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

	message, statusCode, data := controller.UsersService.ProcessGetDetailUser(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *usersControllerImpl) RequestCreateUser(ctx echo.Context) error {
	var request model.RequestCreateUser
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

	message, statusCode, data := controller.UsersService.ProcessCreateUser(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *usersControllerImpl) RequestAllUser(ctx echo.Context) error {
	//Validation
	message, statusCode, data := controller.UsersService.ProcessGetAlllUser(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *usersControllerImpl) RequestUpdateUser(ctx echo.Context) error {
	var request model.RequestUpdateUser
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

	message, statusCode, data := controller.UsersService.ProcessUpdateUser(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *usersControllerImpl) RequestDeleteUser(ctx echo.Context) error {
	var request model.RequestDeleteUser
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

	message, statusCode, data := controller.UsersService.ProcessDeleteUser(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *usersControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
