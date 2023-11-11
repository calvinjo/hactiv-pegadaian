package repository

import (
	"api-service/model"
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type TransactionRepository interface {
	AllLoan(ctx echo.Context, request model.RepoRequestAllLoan) (result model.RepoRespAllLoan)
	DetailLoan(ctx echo.Context, request model.RepoRequestDetailLoan) (result model.RepoRespLoan)
	CreateLoan(ctx echo.Context, request model.RepoRequestCreateLoan) (result model.RepoRespLoan)
	UpdateLoan(ctx echo.Context, request model.RepoRequestUpdateLoan) (result model.RepoRespLoan)
	DeleteLoan(ctx echo.Context, request model.RepoRequestDeleteLoan) (result model.RepoRespLoan)
	UpdateStatusLoan(ctx echo.Context, request model.RepoRequestUpdateStatusLoan) (result model.RepoRespLoan)

	AllInstallment(ctx echo.Context, request model.RepoRequestAllInstallment) (result model.RepoRespAllInstallment)
	DetailInstallment(ctx echo.Context, request model.RepoRequestDetailInstallment) (result model.RepoRespInstallment)
	CreateInstallment(ctx echo.Context, request model.RepoRequestCreateInstallment) (result model.RepoRespInstallment)
	UpdateInstallment(ctx echo.Context, request model.RepoRequestUpdateInstallment) (result model.RepoRespInstallment)
	DeleteInstallment(ctx echo.Context, request model.RepoRequestDeleteInstallment) (result model.RepoRespInstallment)

	AllMaxLoan(ctx echo.Context, request model.RepoRequestAllMaxLoan) (result model.RepoRespAllMaxLoan)
	DetailMaxLoan(ctx echo.Context, request model.RepoRequestDetailMaxLoan) (result model.RepoRespMaxLoan)
	CreateMaxLoan(ctx echo.Context, request model.RepoRequestCreateMaxLoan) (result model.RepoRespMaxLoan)
	UpdateMaxLoan(ctx echo.Context, request model.RepoRequestUpdateMaxLoan) (result model.RepoRespMaxLoan)
	DeleteMaxLoan(ctx echo.Context, request model.RepoRequestDeleteMaxLoan) (result model.RepoRespMaxLoan)
}

type transactionRepositoryImpl struct {
	BaseUrl     string
	RestService *resty.Client
}

func NewTransactionRepository(baseUrl string) TransactionRepository {
	var client = resty.New()
	return &transactionRepositoryImpl{
		BaseUrl:     baseUrl,
		RestService: client,
	}
}

func (repository *transactionRepositoryImpl) AllLoan(ctx echo.Context, request model.RepoRequestAllLoan) (result model.RepoRespAllLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/all-loan")

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

func (repository *transactionRepositoryImpl) DetailLoan(ctx echo.Context, request model.RepoRequestDetailLoan) (result model.RepoRespLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/detail-loan")

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

func (repository *transactionRepositoryImpl) CreateLoan(ctx echo.Context, request model.RepoRequestCreateLoan) (result model.RepoRespLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/create-loan")

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

func (repository *transactionRepositoryImpl) UpdateLoan(ctx echo.Context, request model.RepoRequestUpdateLoan) (result model.RepoRespLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Put(repository.BaseUrl + "/update-loan")

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

func (repository *transactionRepositoryImpl) UpdateStatusLoan(ctx echo.Context, request model.RepoRequestUpdateStatusLoan) (result model.RepoRespLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Patch(repository.BaseUrl + "/update-status-loan")

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

func (repository *transactionRepositoryImpl) DeleteLoan(ctx echo.Context, request model.RepoRequestDeleteLoan) (result model.RepoRespLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Delete(repository.BaseUrl + "/delete-loan")

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

func (repository *transactionRepositoryImpl) AllInstallment(ctx echo.Context, request model.RepoRequestAllInstallment) (result model.RepoRespAllInstallment) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/all-installment")

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

func (repository *transactionRepositoryImpl) DetailInstallment(ctx echo.Context, request model.RepoRequestDetailInstallment) (result model.RepoRespInstallment) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/detail-installment")

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

func (repository *transactionRepositoryImpl) CreateInstallment(ctx echo.Context, request model.RepoRequestCreateInstallment) (result model.RepoRespInstallment) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/create-installment")

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

func (repository *transactionRepositoryImpl) UpdateInstallment(ctx echo.Context, request model.RepoRequestUpdateInstallment) (result model.RepoRespInstallment) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Put(repository.BaseUrl + "/update-installment")

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

func (repository *transactionRepositoryImpl) DeleteInstallment(ctx echo.Context, request model.RepoRequestDeleteInstallment) (result model.RepoRespInstallment) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Delete(repository.BaseUrl + "/delete-installment")

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

func (repository *transactionRepositoryImpl) AllMaxLoan(ctx echo.Context, request model.RepoRequestAllMaxLoan) (result model.RepoRespAllMaxLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/all-max-loan")

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

func (repository *transactionRepositoryImpl) DetailMaxLoan(ctx echo.Context, request model.RepoRequestDetailMaxLoan) (result model.RepoRespMaxLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/detail-max-loan")

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

func (repository *transactionRepositoryImpl) CreateMaxLoan(ctx echo.Context, request model.RepoRequestCreateMaxLoan) (result model.RepoRespMaxLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(repository.BaseUrl + "/create-max-loan")

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

func (repository *transactionRepositoryImpl) UpdateMaxLoan(ctx echo.Context, request model.RepoRequestUpdateMaxLoan) (result model.RepoRespMaxLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Put(repository.BaseUrl + "/update-max-loan")

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

func (repository *transactionRepositoryImpl) DeleteMaxLoan(ctx echo.Context, request model.RepoRequestDeleteMaxLoan) (result model.RepoRespMaxLoan) {
	resp, err := repository.RestService.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Delete(repository.BaseUrl + "/delete-max-loan")

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
