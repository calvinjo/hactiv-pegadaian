package service

import (
	"users-service/config"
	"users-service/model"
	"users-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	ProcessGetAlllUser(ctx echo.Context) (message string, statusCode string, data []model.ResponseAllDataUser)
	ProcessGetDetailUser(ctx echo.Context, request model.RequestGetDetailUser) (message string, statusCode string, data model.ResponseAllDataUser)
	ProcessCreateUser(ctx echo.Context, request model.RequestCreateUser) (message string, statusCode string, data model.ResponseAllDataUser)
	ProcessUpdateUser(ctx echo.Context, request model.RequestUpdateUser) (message string, statusCode string, data model.ResponseAllDataUser)
	ProcessDeleteUser(ctx echo.Context, request model.RequestDeleteUser) (message string, statusCode string, data model.ResponseAllDataUser)
}

type usersServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewUsersService(newUsersReposiory repository.UsersRepository) UsersService {
	return &usersServiceImpl{
		UsersRepository: newUsersReposiory,
	}
}

func (service *usersServiceImpl) ProcessGetAlllUser(ctx echo.Context) (message string, statusCode string, data []model.ResponseAllDataUser) {
	message = ""
	statusCode = config.Success

	result, err := service.UsersRepository.GetAllUser(ctx)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	err = mapstructure.Decode(result, &data)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed
	}

	return
}

func (service *usersServiceImpl) ProcessGetDetailUser(ctx echo.Context, request model.RequestGetDetailUser) (message string, statusCode string, data model.ResponseAllDataUser) {
	message = ""
	statusCode = config.Success

	result, err := service.UsersRepository.GetDetailUser(ctx, request.Filter)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	err = mapstructure.Decode(result, &data)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed
	}

	return
}

func (service *usersServiceImpl) ProcessCreateUser(ctx echo.Context, request model.RequestCreateUser) (message string, statusCode string, data model.ResponseAllDataUser) {
	message = ""
	statusCode = config.Success

	var reqCreateUser model.Users
	err_decode := mapstructure.Decode(request, &reqCreateUser)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	// Crypt pass
	bytesPassword, errCrypt := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if errCrypt != nil {
		message = errCrypt.Error()
		statusCode = config.Failed

		return
	}
	reqCreateUser.Password = string(bytesPassword)

	result, err := service.UsersRepository.CreateUser(ctx, reqCreateUser)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	err = mapstructure.Decode(result, &data)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed
	}

	return
}

func (service *usersServiceImpl) ProcessUpdateUser(ctx echo.Context, request model.RequestUpdateUser) (message string, statusCode string, data model.ResponseAllDataUser) {
	message = ""
	statusCode = config.Success

	var RequestUpdateUser model.UpdateUsers
	err_decode := mapstructure.Decode(request, &RequestUpdateUser)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	// Crypt pass
	bytesPassword, errCrypt := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if errCrypt != nil {
		message = errCrypt.Error()
		statusCode = config.Failed

		return
	}
	RequestUpdateUser.Password = string(bytesPassword)

	result, err := service.UsersRepository.UpdateUser(ctx, request.Filter, RequestUpdateUser)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	err = mapstructure.Decode(result, &data)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed
	}

	return
}

func (service *usersServiceImpl) ProcessDeleteUser(ctx echo.Context, request model.RequestDeleteUser) (message string, statusCode string, data model.ResponseAllDataUser) {
	message = ""
	statusCode = config.Success

	result, err := service.UsersRepository.DeleteUser(ctx, request.Filter)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	err = mapstructure.Decode(result, &data)

	if err != nil {
		message = err.Error()
		statusCode = config.Failed
	}

	return
}
