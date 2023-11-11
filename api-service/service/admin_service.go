package service

import (
	"api-service/config"
	"api-service/model"
	"api-service/repository"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type AdminService interface {
	ProcessGetListUsers(ctx echo.Context) (message string, statusCode string, data []model.ResponseAllDataUsers)
	ProcessGetDetailUsers(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataUsers)
	ProcessCreateUsers(ctx echo.Context, request model.RequestCreateUsers) (message string, statusCode string, data model.ResponseAllDataUsers)
	ProcessUpdateUsers(ctx echo.Context, request model.RequestUpdateUsers, id string) (message string, statusCode string, data model.ResponseAllDataUsers)
	ProcessDeleteUsers(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataUsers)

	ProcessGetListLoan(ctx echo.Context, request model.RequestListLoan) (message string, statusCode string, data []model.ResponseAllDataLoan)
	ProcessGetDetailLoan(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataLoan)
	ProcessUpdateStatusLoan(ctx echo.Context, request model.RequestUpdateStatusLoan, id string) (message string, statusCode string, data model.ResponseAllDataLoan)

	ProcessGetListMaxLoan(ctx echo.Context, request model.RequestListMaxLoan) (message string, statusCode string, data []model.ResponseAllDataMaxLoan)
	ProcessGetDetailMaxLoan(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataMaxLoan)
	ProcessUpdateMaxLoan(ctx echo.Context, request model.RequestUpdateMaxLoan, id string) (message string, statusCode string, data model.ResponseAllDataMaxLoan)

	ProcessGetListInstallment(ctx echo.Context, request model.RequestListInstallment) (message string, statusCode string, data []model.ResponseAllDataInstallment)
}

type adminServiceImpl struct {
	UsersRepository       repository.UsersRepository
	TransactionRepository repository.TransactionRepository
}

func NewAdminService(newUsersRepository repository.UsersRepository, newTransactionRepository repository.TransactionRepository) AdminService {
	return &adminServiceImpl{
		UsersRepository:       newUsersRepository,
		TransactionRepository: newTransactionRepository,
	}
}

func (service *adminServiceImpl) ProcessGetListUsers(ctx echo.Context) (message string, statusCode string, data []model.ResponseAllDataUsers) {
	message = ""
	statusCode = config.Success

	result := service.UsersRepository.AllUsers(ctx)
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

func (service *adminServiceImpl) ProcessGetDetailUsers(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataUsers) {
	message = ""
	statusCode = config.Success

	result := service.UsersRepository.DetailUser(ctx, model.RepoRequestDetailUser{
		Filter: map[string]interface{}{
			"id": id,
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

func (service *adminServiceImpl) ProcessCreateUsers(ctx echo.Context, request model.RequestCreateUsers) (message string, statusCode string, data model.ResponseAllDataUsers) {
	message = ""
	statusCode = config.Success

	var reqCreateUser model.RepoRequestCreateUser
	err := mapstructure.Decode(request, &reqCreateUser)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}
	reqCreateUser.CreatedAt = int(time.Now().Unix())

	result := service.UsersRepository.CreateUser(ctx, reqCreateUser)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

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

func (service *adminServiceImpl) ProcessUpdateUsers(ctx echo.Context, request model.RequestUpdateUsers, id string) (message string, statusCode string, data model.ResponseAllDataUsers) {
	message = ""
	statusCode = config.Success

	var reqUpdateUser model.RepoRequestUpdateUser
	err := mapstructure.Decode(request, &reqUpdateUser)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}
	reqUpdateUser.Filter = map[string]interface{}{
		"id": id,
	}

	result := service.UsersRepository.UpdateUser(ctx, reqUpdateUser)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

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

func (service *adminServiceImpl) ProcessDeleteUsers(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataUsers) {
	message = ""
	statusCode = config.Success

	result := service.UsersRepository.DeleteUser(ctx, model.RepoRequestDeleteUser{
		Filter: map[string]interface{}{
			"id": id,
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

func (service *adminServiceImpl) ProcessGetListLoan(ctx echo.Context, request model.RequestListLoan) (message string, statusCode string, data []model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	dataFilter := make(map[string]interface{})
	if len(request.StatusLoan) > 0 {
		dataFilter["status_loan"] = request.StatusLoan
	}

	if len(request.UserId) > 0 {
		dataFilter["user_id"] = request.UserId
	}

	result := service.TransactionRepository.AllLoan(ctx, model.RepoRequestAllLoan{
		Filter: dataFilter,
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

func (service *adminServiceImpl) ProcessGetDetailLoan(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataLoan) {
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

	errDecode := mapstructure.Decode(result.Data, &data)
	if errDecode != nil {
		message = errDecode.Error()
		statusCode = config.Failed
	}
	return
}

func (service *adminServiceImpl) ProcessUpdateStatusLoan(ctx echo.Context, request model.RequestUpdateStatusLoan, id string) (message string, statusCode string, data model.ResponseAllDataLoan) {
	message = ""
	statusCode = config.Success

	resultDetail := service.TransactionRepository.DetailLoan(ctx, model.RepoRequestDetailLoan{
		Filter: map[string]interface{}{
			"id": id,
		},
	})

	if resultDetail.IsError {
		message = resultDetail.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	if resultDetail.Data.StatusLoan != "pending" {
		message = "Loan Status Not Pending"
		statusCode = config.Failed

		return
	}

	var reqUpdateStatusLoan model.RepoRequestUpdateStatusLoan
	err := mapstructure.Decode(request, &reqUpdateStatusLoan)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}
	reqUpdateStatusLoan.Filter = map[string]interface{}{
		"id": id,
	}

	result := service.TransactionRepository.UpdateStatusLoan(ctx, reqUpdateStatusLoan)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

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

func (service *adminServiceImpl) ProcessGetListMaxLoan(ctx echo.Context, request model.RequestListMaxLoan) (message string, statusCode string, data []model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	dataFilter := make(map[string]interface{})
	if len(request.UserId) > 0 {
		dataFilter["user_id"] = request.UserId
	}

	result := service.TransactionRepository.AllMaxLoan(ctx, model.RepoRequestAllMaxLoan{
		Filter: dataFilter,
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

func (service *adminServiceImpl) ProcessGetDetailMaxLoan(ctx echo.Context, id string) (message string, statusCode string, data model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	result := service.TransactionRepository.DetailMaxLoan(ctx, model.RepoRequestDetailMaxLoan{
		Filter: map[string]interface{}{
			"id": id,
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

func (service *adminServiceImpl) ProcessUpdateMaxLoan(ctx echo.Context, request model.RequestUpdateMaxLoan, id string) (message string, statusCode string, data model.ResponseAllDataMaxLoan) {
	message = ""
	statusCode = config.Success

	var reqUpdateMaxLoan model.RepoRequestUpdateMaxLoan
	err := mapstructure.Decode(request, &reqUpdateMaxLoan)
	if err != nil {
		message = err.Error()
		statusCode = config.Failed

		return
	}
	reqUpdateMaxLoan.Filter = map[string]interface{}{
		"id": id,
	}

	result := service.TransactionRepository.UpdateMaxLoan(ctx, reqUpdateMaxLoan)

	if result.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

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

func (service *adminServiceImpl) ProcessGetListInstallment(ctx echo.Context, request model.RequestListInstallment) (message string, statusCode string, data []model.ResponseAllDataInstallment) {
	message = ""
	statusCode = config.Success

	result := service.TransactionRepository.AllInstallment(ctx, model.RepoRequestAllInstallment{})
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
