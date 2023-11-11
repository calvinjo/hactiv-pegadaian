package service

import (
	"transaction-service/config"
	"transaction-service/model"
	"transaction-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type MaxLoanService interface {
	ProcessGetAllMaxLoan(ctx echo.Context, request model.RequestGetAllMaxLoan) (message string, statusCode string, data []model.ResponseAllDataMaxLoan)
	ProcessGetDetailMaxLoan(ctx echo.Context, request model.RequestGetDetailMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan)
	ProcessCreateMaxLoan(ctx echo.Context, request model.RequestCreateMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan)
	ProcessUpdateMaxLoan(ctx echo.Context, request model.RequestUpdateMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan)
	ProcessDeleteMaxLoan(ctx echo.Context, request model.RequestDeleteMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan)
}

type maxLoanServiceImpl struct {
	MaxLoanRepository repository.MaxLoanRepository
}

func NewMaxLoanService(newMaxLoanRepository repository.MaxLoanRepository) MaxLoanService {
	return &maxLoanServiceImpl{
		MaxLoanRepository: newMaxLoanRepository,
	}
}

func (service *maxLoanServiceImpl) ProcessGetAllMaxLoan(ctx echo.Context, request model.RequestGetAllMaxLoan) (message string, statusCode string, data []model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	result, err := service.MaxLoanRepository.GetAllMaxLoan(ctx, request.Filter)
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

func (service *maxLoanServiceImpl) ProcessGetDetailMaxLoan(ctx echo.Context, request model.RequestGetDetailMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	result, err := service.MaxLoanRepository.GetDetailMaxLoan(ctx, request.Filter)

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

func (service *maxLoanServiceImpl) ProcessCreateMaxLoan(ctx echo.Context, request model.RequestCreateMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	var reqCreateMaxLoan model.MaxLoan
	err_decode := mapstructure.Decode(request, &reqCreateMaxLoan)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	result, err := service.MaxLoanRepository.CreateMaxLoan(ctx, reqCreateMaxLoan)

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

func (service *maxLoanServiceImpl) ProcessUpdateMaxLoan(ctx echo.Context, request model.RequestUpdateMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	var RequestUpdateMaxLoan model.UpdateMaxLoan
	err_decode := mapstructure.Decode(request, &RequestUpdateMaxLoan)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	result, err := service.MaxLoanRepository.UpdateMaxLoan(ctx, request.Filter, RequestUpdateMaxLoan)

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

func (service *maxLoanServiceImpl) ProcessDeleteMaxLoan(ctx echo.Context, request model.RequestDeleteMaxLoan) (message string, statusCode string, data model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	result, err := service.MaxLoanRepository.DeleteMaxLoan(ctx, request.Filter)

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
