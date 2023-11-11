package service

import (
	"transaction-service/config"
	"transaction-service/model"
	"transaction-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type LoanService interface {
	ProcessGetAllLoan(ctx echo.Context, request model.RequestGetAllLoan) (message string, statusCode string, data []model.ResponseAllDataLoan)
	ProcessGetDetailLoan(ctx echo.Context, request model.RequestGetDetailLoan) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessCreateLoan(ctx echo.Context, request model.RequestCreateLoan) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessUpdateLoan(ctx echo.Context, request model.RequestUpdateLoan) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessUpdateStatusLoan(ctx echo.Context, request model.RequestUpdateStatusLoan) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessDeleteLoan(ctx echo.Context, request model.RequestDeleteLoan) (message string, statusCode string, data model.ResponseAllDataLoan)
}

type loanServiceImpl struct {
	LoanRepository    repository.LoanRepository
	MaxLoanRepository repository.MaxLoanRepository
}

func NewLoanService(newLoanRepository repository.LoanRepository, newMaxLoanRepository repository.MaxLoanRepository) LoanService {
	return &loanServiceImpl{
		LoanRepository:    newLoanRepository,
		MaxLoanRepository: newMaxLoanRepository,
	}
}

func (service *loanServiceImpl) ProcessGetAllLoan(ctx echo.Context, request model.RequestGetAllLoan) (message string, statusCode string, data []model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	result, err := service.LoanRepository.GetAllLoan(ctx, request.Filter)
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

func (service *loanServiceImpl) ProcessGetDetailLoan(ctx echo.Context, request model.RequestGetDetailLoan) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	result, err := service.LoanRepository.GetDetailLoan(ctx, request.Filter)

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

func (service *loanServiceImpl) ProcessCreateLoan(ctx echo.Context, request model.RequestCreateLoan) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	var reqCreateLoan model.Loan
	err_decode := mapstructure.Decode(request, &reqCreateLoan)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	isExceedLimit, err := service.CheckMaxLimit(ctx, reqCreateLoan.UserID, reqCreateLoan.Nominal)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	if isExceedLimit {
		message = "Exceed Limit Loan"
		statusCode = config.Exceed

		return
	}

	reqCreateLoan.CostInstallment = service.CalculateInstallment(ctx, reqCreateLoan.Nominal, reqCreateLoan.PeriodInstallment)
	reqCreateLoan.RemainInstallment = reqCreateLoan.Nominal

	result, err := service.LoanRepository.CreateLoan(ctx, reqCreateLoan)

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

func (service *loanServiceImpl) ProcessUpdateLoan(ctx echo.Context, request model.RequestUpdateLoan) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	var RequestUpdateLoan model.UpdateLoan
	err_decode := mapstructure.Decode(request, &RequestUpdateLoan)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	isExceedLimit, err := service.CheckMaxLimit(ctx, RequestUpdateLoan.UserID, RequestUpdateLoan.Nominal)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	if isExceedLimit {
		message = "Exceed Limit Loan"
		statusCode = config.Exceed

		return
	}

	RequestUpdateLoan.CostInstallment = service.CalculateInstallment(ctx, RequestUpdateLoan.Nominal, RequestUpdateLoan.PeriodInstallment)
	RequestUpdateLoan.RemainInstallment = RequestUpdateLoan.Nominal

	result, err := service.LoanRepository.UpdateLoan(ctx, request.Filter, RequestUpdateLoan)

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

func (service *loanServiceImpl) ProcessUpdateStatusLoan(ctx echo.Context, request model.RequestUpdateStatusLoan) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	var RequestUpdateLoan model.UpdateLoan
	err_decode := mapstructure.Decode(request, &RequestUpdateLoan)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	result, err := service.LoanRepository.UpdateLoan(ctx, request.Filter, RequestUpdateLoan)

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

func (service *loanServiceImpl) ProcessDeleteLoan(ctx echo.Context, request model.RequestDeleteLoan) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	result, err := service.LoanRepository.DeleteLoan(ctx, request.Filter)

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

func (service *loanServiceImpl) CheckMaxLimit(ctx echo.Context, userId int64, nominal int64) (isExceedLimit bool, err error) {
	result, err := service.MaxLoanRepository.GetDetailMaxLoan(ctx, map[string]interface{}{
		"user_id": userId,
	})
	if err != nil {
		return
	}

	//Jika belum ada limit
	if (result == model.MaxLoan{}) {
		if nominal > config.MaxLimit {
			isExceedLimit = true
			return
		}

		_, err = service.MaxLoanRepository.CreateMaxLoan(ctx, model.MaxLoan{
			UserID:  userId,
			Current: nominal,
			Limit:   config.MaxLimit,
		})

		return
	}

	if (result.Current + nominal) > result.Limit {
		isExceedLimit = true
		return
	}

	//Incerease current limit
	_, err = service.MaxLoanRepository.UpdateMaxLoan(ctx, map[string]interface{}{
		"user_id": userId,
	}, model.UpdateMaxLoan{
		Current: result.Current + nominal,
	})

	return
}

func (service *loanServiceImpl) CalculateInstallment(ctx echo.Context, nominal int64, installmentPeriod int64) (costsInstallment int64) {
	return nominal / installmentPeriod
}
