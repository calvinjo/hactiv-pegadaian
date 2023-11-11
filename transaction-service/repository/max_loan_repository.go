package repository

import (
	"errors"
	"transaction-service/config"
	"transaction-service/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MaxLoanRepository interface {
	GetAllMaxLoan(ctx echo.Context, filter map[string]interface{}) (result []model.MaxLoan, err error)
	GetDetailMaxLoan(ctx echo.Context, filter map[string]interface{}) (result model.MaxLoan, err error)
	CreateMaxLoan(ctx echo.Context, request model.MaxLoan) (result model.MaxLoan, err error)
	UpdateMaxLoan(ctx echo.Context, filter map[string]interface{}, request model.UpdateMaxLoan) (result model.UpdateMaxLoan, err error)
	DeleteMaxLoan(ctx echo.Context, filter map[string]interface{}) (result model.MaxLoan, err error)
}

type MaxLoanRepositoryImpl struct {
	Postgre config.PostgreSql
}

func NewMaxLoanRepository(db config.PostgreSql) MaxLoanRepository {
	return &MaxLoanRepositoryImpl{
		Postgre: db,
	}
}

func (repository *MaxLoanRepositoryImpl) GetAllMaxLoan(ctx echo.Context, filter map[string]interface{}) (result []model.MaxLoan, err error) {
	tx := repository.Postgre.Db.Where(filter).Find(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (repository *MaxLoanRepositoryImpl) GetDetailMaxLoan(ctx echo.Context, filter map[string]interface{}) (result model.MaxLoan, err error) {
	tx := repository.Postgre.Db.Where(filter).First(&result)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			err = tx.Error
			return
		}
	}
	return
}

func (repository *MaxLoanRepositoryImpl) CreateMaxLoan(ctx echo.Context, request model.MaxLoan) (result model.MaxLoan, err error) {
	tx := repository.Postgre.Db.Create(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *MaxLoanRepositoryImpl) UpdateMaxLoan(ctx echo.Context, filter map[string]interface{}, request model.UpdateMaxLoan) (result model.UpdateMaxLoan, err error) {
	tx := repository.Postgre.Db.Model(model.UpdateMaxLoan{}).Where(filter).Table("max_loans").Updates(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *MaxLoanRepositoryImpl) DeleteMaxLoan(ctx echo.Context, filter map[string]interface{}) (result model.MaxLoan, err error) {
	tx := repository.Postgre.Db.Where(filter).Delete(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
