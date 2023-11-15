package repository

import (
	"api-service/model"
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type UsersRepository interface {
	AllUsers(ctx echo.Context) (result model.RepoRespAllUsers)
	DetailUser(ctx echo.Context, request model.RepoRequestDetailUser) (result model.RepoRespUsers)
	CreateUser(ctx echo.Context, request model.RepoRequestCreateUser) (result model.RepoRespUsers)
	UpdateUser(ctx echo.Context, request model.RepoRequestUpdateUser) (result model.RepoRespUsers)
	DeleteUser(ctx echo.Context, request model.RepoRequestDeleteUser) (result model.RepoRespUsers)
}

type usersRepositoryImpl struct {
	BaseUrl     string
	RestService *resty.Client
}

func NewUsersRepository(baseUrl string) UsersRepository {
	var client = resty.New()
	return &usersRepositoryImpl{
		BaseUrl:     baseUrl,
		RestService: client,
	}
}

func (repository *usersRepositoryImpl) AllUsers(ctx echo.Context) (result model.RepoRespAllUsers) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		Get(repository.BaseUrl + "/all-user")

	if err != nil {
		result.IsError = true
		result.ErrorMessage = err

		return
	}
	if resp.StatusCode() == 404 {
		result.IsNotFound = true
		return
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 422 && resp.StatusCode() != 429 {
		result.IsError = true
		result.ErrorMessage = errors.New("Response not 200")
		return
	}

	if parseErr := json.Unmarshal(resp.Body(), &result); parseErr != nil {
		result.IsError = true
		result.ErrorMessage = parseErr
		return
	}

	return
}

func (repository *usersRepositoryImpl) DetailUser(ctx echo.Context, request model.RepoRequestDetailUser) (result model.RepoRespUsers) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/detail-user")

	if err != nil {
		result.IsError = true
		result.ErrorMessage = err

		return
	}
	if resp.StatusCode() == 404 {
		result.IsNotFound = true
		return
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 422 && resp.StatusCode() != 429 {
		result.IsError = true
		result.ErrorMessage = errors.New("Response not 200")
		return
	}

	if parseErr := json.Unmarshal(resp.Body(), &result); parseErr != nil {
		result.IsError = true
		result.ErrorMessage = parseErr
		return
	}

	return
}

func (repository *usersRepositoryImpl) CreateUser(ctx echo.Context, request model.RepoRequestCreateUser) (result model.RepoRespUsers) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/create-user")

	if err != nil {
		result.IsError = true
		result.ErrorMessage = err

		return
	}

	if resp.StatusCode() == 404 {
		result.IsNotFound = true
		return
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 422 && resp.StatusCode() != 429 {
		result.IsError = true
		result.ErrorMessage = errors.New("Response not 200")
		return
	}

	if parseErr := json.Unmarshal(resp.Body(), &result); parseErr != nil {
		result.IsError = true
		result.ErrorMessage = parseErr
		return
	}

	return
}

func (repository *usersRepositoryImpl) UpdateUser(ctx echo.Context, request model.RepoRequestUpdateUser) (result model.RepoRespUsers) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Put(repository.BaseUrl + "/update-user")

	if err != nil {
		result.IsError = true
		result.ErrorMessage = err

		return
	}
	if resp.StatusCode() == 404 {
		result.IsNotFound = true
		return
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 422 && resp.StatusCode() != 429 {
		result.IsError = true
		result.ErrorMessage = errors.New("Response not 200")
		return
	}

	if parseErr := json.Unmarshal(resp.Body(), &result); parseErr != nil {
		result.IsError = true
		result.ErrorMessage = parseErr
		return
	}

	return
}

func (repository *usersRepositoryImpl) DeleteUser(ctx echo.Context, request model.RepoRequestDeleteUser) (result model.RepoRespUsers) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Delete(repository.BaseUrl + "/delete-user")

	if err != nil {
		result.IsError = true
		result.ErrorMessage = err

		return
	}
	if resp.StatusCode() == 404 {
		result.IsNotFound = true
		return
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 422 && resp.StatusCode() != 429 {
		result.IsError = true
		result.ErrorMessage = errors.New("Response not 200")
		return
	}

	if parseErr := json.Unmarshal(resp.Body(), &result); parseErr != nil {
		result.IsError = true
		result.ErrorMessage = parseErr
		return
	}

	return
}
