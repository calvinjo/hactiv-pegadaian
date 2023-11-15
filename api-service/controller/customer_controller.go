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

type CustomerController interface {
	RequestGetProfile(ctx echo.Context) error
	RequestCreateLoan(ctx echo.Context) error
	RequestGetListLoan(ctx echo.Context) error
	RequestUpdateLoan(ctx echo.Context) error
	RequestCreateInstallment(ctx echo.Context) error
	RequestGetHistoryInstallment(ctx echo.Context) error
	RequestGetDetailLoan(ctx echo.Context) error
}

type customerControllerImpl struct {
	CustomerService service.CustomerService
	Validator       *validator.Validate
}

func NewCustomerController(newCustomerService service.CustomerService, newValidator *validator.Validate) CustomerController {
	return &customerControllerImpl{
		CustomerService: newCustomerService,
		Validator:       newValidator,
	}
}

// @Summary Customer - Get Profile
// @ID customer-get-profile
// @Accept json
// @Produce json
// @Router /customer/profile [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagGetProfile "Success"
// @Failure 404 {object} model.Response "User not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestGetProfile(ctx echo.Context) error {

	message, statusCode, data := controller.CustomerService.ProcessGetProfile(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Customer - Create Loan
// @ID customer-create-loan
// @Accept json
// @Produce json
// @Router /customer/create-loan [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestCreateLoan true "Request"
// @Success 200 {object} model.ResponseSwagDataLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestCreateLoan(ctx echo.Context) error {

	var request model.RequestCreateLoan
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

	message, statusCode, data := controller.CustomerService.ProcessCreateLoan(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Customer - List Loan
// @ID customer-list-loan
// @Accept json
// @Produce json
// @Router /customer/list-loan [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param status query string true "Status value (pending, paid, approved, disapproved)" Enums(pending, paid, approved, disapproved)
// @Success 200 {object} model.ResponseSwagDataAllLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestGetListLoan(ctx echo.Context) error {

	message, statusCode, data := controller.CustomerService.ProcessGetListLoan(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Customer - Update Loan
// @ID customer-update-loan
// @Accept json
// @Produce json
// @Router /customer/update-loan/{id} [put]
// @Security ApiKeyAuth
// @Param id path int true "Loan ID" Format(int64)
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestUpdateLoan true "Request"
// @Success 200 {object} model.ResponseSwagDataLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestUpdateLoan(ctx echo.Context) error {
	id := ctx.Param("id")

	var request model.RequestUpdateLoan
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

	message, statusCode, data := controller.CustomerService.ProcessUpdateLoan(ctx, request, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Customer - Create Installment
// @ID customer-create-installment
// @Accept json
// @Produce json
// @Router /customer/create-installment [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param Request body model.RequestCreateInstallment true "Request"
// @Success 200 {object} model.ResponseSwagDataInstallment "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestCreateInstallment(ctx echo.Context) error {
	var request model.RequestCreateInstallment
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

	message, statusCode, data := controller.CustomerService.ProcessCreateInstallment(ctx, request)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Customer - History Installment
// @ID customer-history-installment
// @Accept json
// @Produce json
// @Router /customer/history-installment [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Success 200 {object} model.ResponseSwagDataAllInstallment "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestGetHistoryInstallment(ctx echo.Context) error {

	message, statusCode, data := controller.CustomerService.ProcessGetHistoryInstallment(ctx)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

// @Summary Customer - Detail Loan
// @ID customer-detail-loan
// @Accept json
// @Produce json
// @Router /customer/detail-loan/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer token"
// @Param id path int true "Loan ID" Format(int64)
// @Success 200 {object} model.ResponseSwagDetailLoan "Success"
// @Failure 404 {object} model.Response "Not found"
// @Failure 412 {object} model.Response "Failed"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Tags Customer
func (controller *customerControllerImpl) RequestGetDetailLoan(ctx echo.Context) error {

	id := ctx.Param("id")

	message, statusCode, data := controller.CustomerService.ProcessGetDetailLoan(ctx, id)

	return helpers.GenerateResponse(ctx, statusCode, message, data, nil)
}

func (controller *customerControllerImpl) validation(ctx echo.Context, request interface{}) (bool, error) {
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
