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

type AuthController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type authControllerImpl struct {
	AuthService service.AuthService
	Validator   *validator.Validate
}

func NewAuthController(newAuthService service.AuthService, newValidator *validator.Validate) AuthController {
	return &authControllerImpl{
		AuthService: newAuthService,
		Validator:   newValidator,
	}
}

func (controller *authControllerImpl) Register(ctx echo.Context) error {
	var request model.RequestRegister
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

	message, statusCode, data := controller.AuthService.ProcessRegister(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *authControllerImpl) Login(ctx echo.Context) error {
	var request model.RequestLogin
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

	message, statusCode, data := controller.AuthService.ProcessLogin(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *authControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
