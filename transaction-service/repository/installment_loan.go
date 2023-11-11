package repository

import (
	"transaction-service/config"
	"transaction-service/model"

	"github.com/labstack/echo/v4"
)

type InstallmentRepository interface {
	GetAllInstallment(ctx echo.Context, filter map[string]interface{}) (result []model.Installment, err error)
	GetDetailInstallment(ctx echo.Context, filter map[string]interface{}) (result model.Installment, err error)
	CreateInstallment(ctx echo.Context, request model.Installment) (result model.Installment, err error)
	UpdateInstallment(ctx echo.Context, filter map[string]interface{}, request model.UpdateInstallment) (result model.UpdateInstallment, err error)
	DeleteInstallment(ctx echo.Context, filter map[string]interface{}) (result model.Installment, err error)
}

type InstallmentRepositoryImpl struct {
	Postgre config.PostgreSql
}

func NewInstallmentRepository(db config.PostgreSql) InstallmentRepository {
	return &InstallmentRepositoryImpl{
		Postgre: db,
	}
}

func (repository *InstallmentRepositoryImpl) GetAllInstallment(ctx echo.Context, filter map[string]interface{}) (result []model.Installment, err error) {
	tx := repository.Postgre.Db.Where(filter).Find(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (repository *InstallmentRepositoryImpl) GetDetailInstallment(ctx echo.Context, filter map[string]interface{}) (result model.Installment, err error) {
	tx := repository.Postgre.Db.Where(filter).First(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (repository *InstallmentRepositoryImpl) CreateInstallment(ctx echo.Context, request model.Installment) (result model.Installment, err error) {
	tx := repository.Postgre.Db.Create(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *InstallmentRepositoryImpl) UpdateInstallment(ctx echo.Context, filter map[string]interface{}, request model.UpdateInstallment) (result model.UpdateInstallment, err error) {
	tx := repository.Postgre.Db.Model(model.UpdateInstallment{}).Where(filter).Table("installments").Updates(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *InstallmentRepositoryImpl) DeleteInstallment(ctx echo.Context, filter map[string]interface{}) (result model.Installment, err error) {
	tx := repository.Postgre.Db.Where(filter).Delete(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
