package service

import (
	"api-service/config"
	"api-service/model"
	"api-service/repository"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type CustomerService interface {
	ProcessGetProfile(ctx echo.Context) (message string, statusCode string, data model.ResponseAllDataUsers)
	ProcessCreateLoan(ctx echo.Context, request model.RequestCreateLoan) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessGetListLoan(ctx echo.Context) (message string, statusCode string, data model.ResponseListLoanCustomer)
	ProcessUpdateLoan(ctx echo.Context, request model.RequestUpdateLoan, id string) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessCreateInstallment(ctx echo.Context, request model.RequestCreateInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment)
	ProcessGetHistoryInstallment(ctx echo.Context) (message string, statusCode string, data []model.ResponseAllDataInstallment)
	ProcessGetDetailLoan(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseDetailLoan)
}

type customerServiceImpl struct {
	UsersRepository       repository.UsersRepository
	TransactionRepository repository.TransactionRepository
}

func NewCustomerService(newUsersRepository repository.UsersRepository, newTransactionRepository repository.TransactionRepository) CustomerService {
	return &customerServiceImpl{
		UsersRepository:       newUsersRepository,
		TransactionRepository: newTransactionRepository,
	}
}

func (service *customerServiceImpl) ProcessGetProfile(ctx echo.Context) (message string, statusCode string, data model.ResponseAllDataUsers) {
	message = ""
	statusCode = config.Success

	result := service.UsersRepository.DetailUser(ctx, model.RepoRequestDetailUser{
		Filter: map[string]interface{}{
			"id": ctx.Get("userId"),
		},
	})

	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}
	return
}

func (service *customerServiceImpl) ProcessCreateLoan(ctx echo.Context, request model.RequestCreateLoan) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	var reqCreateLoan model.RepoRequestCreateLoan
	err := mapstructure.Decode(request, &reqCreateLoan)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		message = "Value is not of type int64"
		statusCode = config.Failed
		return
	}
	reqCreateLoan.UserID = userId
	reqCreateLoan.StatusLoan = "pending"
	reqCreateLoan.CreatedAt = int(time.Now().Unix())

	result := service.TransactionRepository.CreateLoan(ctx, reqCreateLoan)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	// Error limit amount
	if result.ResponseCode == "429" {
		message = result.Message
		statusCode = config.Failed

		return
	}

	// Error validate data loan
	if result.ResponseCode == "422" {
		message = result.Message
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}
	return
}

func (service *customerServiceImpl) ProcessGetListLoan(ctx echo.Context) (message string, statusCode string, data model.ResponseListLoanCustomer) {
	message = ""
	statusCode = config.Success

	filterStatus := ctx.QueryParam("status")
	filterLoan := map[string]interface{}{"user_id": ctx.Get("userId")}
	if len(filterStatus) > 0 && (filterStatus == "approved" || filterStatus == "pending" || filterStatus == "disapproved" || filterStatus == "paid") {
		filterLoan = map[string]interface{}{"user_id": ctx.Get("userId"), "status_loan": filterStatus}
	}

	result := service.TransactionRepository.AllLoan(ctx, model.RepoRequestAllLoan{
		Filter: filterLoan,
	})
	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data.DataLoan)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}

	data.TotalInstallment = 0
	data.TotalInstallmentPaid = 0
	data.TotalInstallmentRemaining = 0
	for _, value := range result.Data {
		data.TotalInstallment += value.Nominal
		data.TotalInstallmentPaid += value.Paid
		data.TotalInstallmentRemaining += value.RemainInstallment
	}
	return
}

func (service *customerServiceImpl) ProcessUpdateLoan(ctx echo.Context, request model.RequestUpdateLoan, id string) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	resultDetail := service.TransactionRepository.DetailLoan(ctx, model.RepoRequestDetailLoan{
		Filter: map[string]interface{}{
			"id": id,
		},
	})

	if resultDetail.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

	if resultDetail.IsError {
		message = resultDetail.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	//Update hanya boleh dilakukan untuk status loan yang sedang pending
	if resultDetail.Data.StatusLoan != "pending" {
		message = "The Loan Has Been Processed"
		statusCode = config.Failed

		return
	}

	var reqUpdateLoan model.RepoRequestUpdateLoan
	err := mapstructure.Decode(request, &reqUpdateLoan)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}
	reqUpdateLoan.StatusLoan = "pending"

	reqUpdateLoan.Filter = map[string]interface{}{
		"id": id,
	}

	result := service.TransactionRepository.UpdateLoan(ctx, reqUpdateLoan)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	// Error limit amount
	if result.ResponseCode == "429" {
		message = result.Message
		statusCode = config.Failed

		return
	}

	// Error validate data loan
	if result.ResponseCode == "422" {
		message = result.Message
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}
	return
}

func (service *customerServiceImpl) ProcessCreateInstallment(ctx echo.Context, request model.RequestCreateInstallment) (message string, statusCode string, data model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	var reqCreateInstallment model.RepoRequestCreateInstallment
	err := mapstructure.Decode(request, &reqCreateInstallment)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}

	resultDetailLoan := service.TransactionRepository.DetailLoan(ctx, model.RepoRequestDetailLoan{
		Filter: map[string]interface{}{
			"id": request.LoanID,
		},
	})

	if resultDetailLoan.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

	if resultDetailLoan.IsError {
		message = resultDetailLoan.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	//Pembayaran hanya boleh untuk data loan yang sudah aktif
	if resultDetailLoan.Data.StatusLoan != "approved" {
		message = "The Loan Is Not Active"
		statusCode = config.Failed

		return
	}

	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		message = "Value is not of type int64"
		statusCode = config.Failed
		return
	}
	reqCreateInstallment.UserID = userId
	reqCreateInstallment.PaymentAt = int(time.Now().Unix())

	result := service.TransactionRepository.CreateInstallment(ctx, reqCreateInstallment)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	// Error limit amount
	if result.ResponseCode == "429" {
		message = result.Message
		statusCode = config.Failed

		return
	}

	// Error validate data loan
	if result.ResponseCode == "422" {
		message = result.Message
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}
	return
}

func (service *customerServiceImpl) ProcessGetHistoryInstallment(ctx echo.Context) (message string, statusCode string, data []model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	result := service.TransactionRepository.AllInstallment(ctx, model.RepoRequestAllInstallment{
		Filter: map[string]interface{}{"user_id": ctx.Get("userId")},
	})

	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}
	return
}

func (service *customerServiceImpl) ProcessGetDetailLoan(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseDetailLoan) {
	message = ""
	statusCode = config.Success

	result := service.TransactionRepository.DetailLoan(ctx, model.RepoRequestDetailLoan{
		Filter: map[string]interface{}{
			"id": id,
		},
	})

	if result.IsError {
		message = result.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	errDecode := mapstructure.Decode(result.Data, &data.DataLoan)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}

	resultInstallment := service.TransactionRepository.AllInstallment(ctx, model.RepoRequestAllInstallment{
		Filter: map[string]interface{}{
			"loan_id": id,
			"user_id": result.Data.UserID,
		},
	})

	if resultInstallment.IsError {
		message = resultInstallment.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	errDecode = mapstructure.Decode(resultInstallment.Data, &data.HistoryInstallment)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}

	return
}
