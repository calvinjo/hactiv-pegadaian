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

type AdminController interface {
	RequestGetListUsers(ctx echo.Context) error
	RequestGetDetailUsers(ctx echo.Context) error
	RequestCreateUsers(ctx echo.Context) error
	RequestUpdateUsers(ctx echo.Context) error
	RequestDeleteUsers(ctx echo.Context) error

	RequestGetListLoan(ctx echo.Context) error
	RequestGetDetailLoan(ctx echo.Context) error
	RequestUpdateStatusLoan(ctx echo.Context) error

	RequestGetListMaxLoan(ctx echo.Context) error
	RequestGetDetailMaxLoan(ctx echo.Context) error
	RequestUpdateMaxLoan(ctx echo.Context) error

	RequestGetListInstallment(ctx echo.Context) error
}

type adminControllerImpl struct {
	AdminService service.AdminService
	Validator    *validator.Validate
}

func NewAdminController(newAdminService service.AdminService, newValidator *validator.Validate) AdminController {
	return &adminControllerImpl{
		AdminService: newAdminService,
		Validator:    newValidator,
	}
}

// @Summary Admin - List User
// @ID admin-list-user
// @Accept json
// @Produce json
// @Router /admin/list-user [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagDataAllUsers "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetListUsers(ctx echo.Context) error {

	message, statusCode, data := controller.AdminService.ProcessGetListUsers(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Detail User
// @ID admin-detail-user
// @Accept json
// @Produce json
// @Router /admin/detail-user/{id} [get]
// @Param id path int true "User ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagDataUsers "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetDetailUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessGetDetailUsers(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Create User
// @ID admin-create-user
// @Accept json
// @Produce json
// @Router /admin/create-user [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestCreateUsers true "Request"
// @Success 200 {object} model.ResponseSwagDataUsers "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestCreateUsers(ctx echo.Context) error {
	var request model.RequestCreateUsers
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

	message, statusCode, data := controller.AdminService.ProcessCreateUsers(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Update User
// @ID admin-update-user
// @Accept json
// @Produce json
// @Router /admin/update-user/{id} [put]
// @Param id path int true "User ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestUpdateUsers true "Request"
// @Success 200 {object} model.ResponseSwagDataUsers "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestUpdateUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	var request model.RequestUpdateUsers
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

	message, statusCode, data := controller.AdminService.ProcessUpdateUsers(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Delete User
// @ID admin-delete-user
// @Accept json
// @Produce json
// @Router /admin/delete-user/{id} [delete]
// @Param id path int true "User ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagDataUsers "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestDeleteUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessDeleteUsers(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - List Loan
// @ID admin-list-loan
// @Accept json
// @Produce json
// @Router /admin/list-loan [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Params query model.RequestListLoan true "Request"
// @Success 200 {object} model.ResponseSwagDataAllLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetListLoan(ctx echo.Context) error {
	var request model.RequestListLoan
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

	message, statusCode, data := controller.AdminService.ProcessGetListLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Detail Loan
// @ID admin-detail-loan
// @Accept json
// @Produce json
// @Router /admin/detail-loan/{id} [get]
// @Param id path int true "Loan ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagDataLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetDetailLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessGetDetailLoan(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Update Status
// @ID admin-update-status-loan
// @Accept json
// @Produce json
// @Router /admin/update-status-loan/{id} [patch]
// @Param id path int true "Loan ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestUpdateStatusLoan true "Request"
// @Success 200 {object} model.ResponseSwagDataUsers "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestUpdateStatusLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	var request model.RequestUpdateStatusLoan
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

	message, statusCode, data := controller.AdminService.ProcessUpdateStatusLoan(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - List Max Loan
// @ID admin-list-max-loan
// @Accept json
// @Produce json
// @Router /admin/list-max-loan [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Params query model.RequestListMaxLoan true "Request"
// @Success 200 {object} model.ResponseSwagDataAllMaxLoan"Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetListMaxLoan(ctx echo.Context) error {
	var request model.RequestListMaxLoan
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

	message, statusCode, data := controller.AdminService.ProcessGetListMaxLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Detail Max Loan
// @ID admin-detail-max-loan
// @Accept json
// @Produce json
// @Router /admin/detail-max-loan/{id} [get]
// @Param id path int true "Loan ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagDataMaxLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetDetailMaxLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	message, statusCode, data := controller.AdminService.ProcessGetDetailMaxLoan(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - Update Max Loan
// @ID admin-update-max-loan
// @Accept json
// @Produce json
// @Router /admin/update-max-loan/{id} [put]
// @Param id path int true "User ID" Format(int64)
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestUpdateMaxLoan true "Request"
// @Success 200 {object} model.ResponseSwagDataLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestUpdateMaxLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	var request model.RequestUpdateMaxLoan
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

	message, statusCode, data := controller.AdminService.ProcessUpdateMaxLoan(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Admin - List Installment
// @ID admin-list-installment
// @Accept json
// @Produce json
// @Router /admin/list-installment [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Params query model.RequestListInstallment true "Request"
// @Success 200 {object} model.ResponseSwagDataAllInstallment "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Admin
func (controller *adminControllerImpl) RequestGetListInstallment(ctx echo.Context) error {

	var request model.RequestListInstallment
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

	message, statusCode, data := controller.AdminService.ProcessGetListInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *adminControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
