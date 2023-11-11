package service

import (
	"transaction-service/config"
	"transaction-service/model"
	"transaction-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type InstallmentService interface {
	ProcessGetAllInstallment(ctx echo.Context, request model.RequestGetAllInstallment) (message string, statusCode string, data []model.ResponseAllDataInstallment)
	ProcessGetDetailInstallment(ctx echo.Context, request model.RequestGetDetailInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment)
	ProcessCreateInstallment(ctx echo.Context, request model.RequestCreateInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment)
	ProcessUpdateInstallment(ctx echo.Context, request model.RequestUpdateInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment)
	ProcessDeleteInstallment(ctx echo.Context, request model.RequestDeleteInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment)
}

type installmentServiceImpl struct {
	InstallmentRepository repository.InstallmentRepository
	LoanRepository        repository.LoanRepository
	MaxLoanRepository     repository.MaxLoanRepository
}

func NewInstallmentService(newInstallmentRepository repository.InstallmentRepository, newLoanRepository repository.LoanRepository, newMaxLoanRepository repository.MaxLoanRepository) InstallmentService {
	return &installmentServiceImpl{
		InstallmentRepository: newInstallmentRepository,
		LoanRepository:        newLoanRepository,
		MaxLoanRepository:     newMaxLoanRepository,
	}
}

func (service *installmentServiceImpl) ProcessGetAllInstallment(ctx echo.Context, request model.RequestGetAllInstallment) (message string, statusCode string, data []model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	result, err := service.InstallmentRepository.GetAllInstallment(ctx, request.Filter)
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

func (service *installmentServiceImpl) ProcessGetDetailInstallment(ctx echo.Context, request model.RequestGetDetailInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	result, err := service.InstallmentRepository.GetDetailInstallment(ctx, request.Filter)

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

func (service *installmentServiceImpl) ProcessCreateInstallment(ctx echo.Context, request model.RequestCreateInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	var reqCreateInstallment model.Installment
	err_decode := mapstructure.Decode(request, &reqCreateInstallment)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	message, isWaring, err := service.UpdateLoan(ctx, reqCreateInstallment.UserID, reqCreateInstallment.LoanID, reqCreateInstallment.Nominal)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	if isWaring {
		statusCode = config.Invalid
		return
	}

	result, err := service.InstallmentRepository.CreateInstallment(ctx, reqCreateInstallment)
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

func (service *installmentServiceImpl) ProcessUpdateInstallment(ctx echo.Context, request model.RequestUpdateInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	var RequestUpdateInstallment model.UpdateInstallment
	err_decode := mapstructure.Decode(request, &RequestUpdateInstallment)
	if err_decode != nil {
		message = err_decode.Error()
		statusCode = config.Failed

		return
	}

	result, err := service.InstallmentRepository.UpdateInstallment(ctx, request.Filter, RequestUpdateInstallment)

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

func (service *installmentServiceImpl) ProcessDeleteInstallment(ctx echo.Context, request model.RequestDeleteInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	result, err := service.InstallmentRepository.DeleteInstallment(ctx, request.Filter)

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

func (service *installmentServiceImpl) UpdateLoan(ctx echo.Context, userId int64, loanId int64, nominal int64) (message string, isWarning bool, err error) {
	result, err := service.LoanRepository.GetDetailLoan(ctx, map[string]interface{}{
		"user_id": userId,
		"id":      loanId,
	})

	if err != nil {
		message = err.Error()
		return
	}

	if (result == model.Loan{}) {
		message = "Loan Not Found"
		isWarning = true
		return
	}

	//Data installment hanya boleh di lakukan untuk status approved
	if result.StatusLoan != "approved" {
		message = "Loan Is Pending"
		isWarning = true
		return
	}

	if nominal < result.CostInstallment {
		message = "Less Nominal Installments"
		isWarning = true
		return
	}

	if nominal > result.RemainInstallment {
		message = "Nominal Amount Exceeds The Remaining Payment Amount"
		isWarning = true
		return
	}

	remainInstallment := result.Nominal - (result.Paid + nominal)

	var requestUpdateLoan model.UpdateLoan

	//Jika lunas
	if remainInstallment == 0 {
		requestUpdateLoan = model.UpdateLoan{
			RemainInstallment: -1,
			Paid:              result.Paid + nominal,
			StatusLoan:        "paid",
		}

		//Deduct current limit
		resultLimit, errLimit := service.MaxLoanRepository.GetDetailMaxLoan(ctx, map[string]interface{}{
			"user_id": userId,
		})
		if errLimit != nil {
			err = errLimit
			return
		}

		_, errDeductLimit := service.MaxLoanRepository.UpdateMaxLoan(ctx, map[string]interface{}{
			"user_id": userId,
		}, model.UpdateMaxLoan{
			Current: resultLimit.Current - result.Nominal,
		})
		if errDeductLimit != nil {
			err = errDeductLimit
			return
		}

	} else {
		requestUpdateLoan = model.UpdateLoan{
			RemainInstallment: result.Nominal - (result.Paid + nominal),
			Paid:              result.Paid + nominal,
		}
	}

	_, err = service.LoanRepository.UpdateLoan(ctx, map[string]interface{}{
		"user_id": userId,
		"id":      loanId,
	}, requestUpdateLoan)

	return
}
