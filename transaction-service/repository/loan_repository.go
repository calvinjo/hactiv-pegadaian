package repository

import (
	"errors"
	"transaction-service/config"
	"transaction-service/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoanRepository interface {
	GetAllLoan(ctx echo.Context, filter map[string]interface{}) (result []model.Loan, err error)
	GetDetailLoan(ctx echo.Context, filter map[string]interface{}) (result model.Loan, err error)
	CreateLoan(ctx echo.Context, request model.Loan) (result model.Loan, err error)
	UpdateLoan(ctx echo.Context, filter map[string]interface{}, request model.UpdateLoan) (result model.UpdateLoan, err error)
	DeleteLoan(ctx echo.Context, filter map[string]interface{}) (result model.Loan, err error)
}

type LoanRepositoryImpl struct {
	Postgre config.PostgreSql
}

func NewLoanRepository(db config.PostgreSql) LoanRepository {
	return &LoanRepositoryImpl{
		Postgre: db,
	}
}

func (repository *LoanRepositoryImpl) GetAllLoan(ctx echo.Context, filter map[string]interface{}) (result []model.Loan, err error) {
	tx := repository.Postgre.Db.Where(filter).Find(&result)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			err = tx.Error
			return
		}
	}
	return
}

func (repository *LoanRepositoryImpl) GetDetailLoan(ctx echo.Context, filter map[string]interface{}) (result model.Loan, err error) {
	tx := repository.Postgre.Db.Where(filter).First(&result)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			err = tx.Error
			return
		}
	}
	return
}

func (repository *LoanRepositoryImpl) CreateLoan(ctx echo.Context, request model.Loan) (result model.Loan, err error) {
	tx := repository.Postgre.Db.Create(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *LoanRepositoryImpl) UpdateLoan(ctx echo.Context, filter map[string]interface{}, request model.UpdateLoan) (result model.UpdateLoan, err error) {
	tx := repository.Postgre.Db.Model(model.UpdateLoan{}).Where(filter).Table("loans").Updates(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *LoanRepositoryImpl) DeleteLoan(ctx echo.Context, filter map[string]interface{}) (result model.Loan, err error) {
	tx := repository.Postgre.Db.Where(filter).Delete(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
