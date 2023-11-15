package repository

import (
	"errors"
	"users-service/config"
	"users-service/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUser(ctx echo.Context) (result []model.Users, err error)
	GetDetailUser(ctx echo.Context, filter map[string]interface{}) (result model.Users, err error)
	CreateUser(ctx echo.Context, request model.Users) (result model.Users, err error)
	UpdateUser(ctx echo.Context, filter map[string]interface{}, request model.UpdateUsers) (result model.UpdateUsers, err error)
	DeleteUser(ctx echo.Context, filter map[string]interface{}) (result model.Users, err error)
}

type usersRepositoryImpl struct {
	Postgre config.PostgreSql
}

func NewUsersRepository(db config.PostgreSql) UsersRepository {
	return &usersRepositoryImpl{
		Postgre: db,
	}
}

func (repository *usersRepositoryImpl) GetAllUser(ctx echo.Context) (result []model.Users, err error) {
	tx := repository.Postgre.Db.Find(&result)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			err = tx.Error
			return
		}
	}
	return
}

func (repository *usersRepositoryImpl) GetDetailUser(ctx echo.Context, filter map[string]interface{}) (result model.Users, err error) {
	tx := repository.Postgre.Db.Where(filter).First(&result)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			err = tx.Error
			return
		}
	}
	return
}

func (repository *usersRepositoryImpl) CreateUser(ctx echo.Context, request model.Users) (result model.Users, err error) {
	tx := repository.Postgre.Db.Create(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *usersRepositoryImpl) UpdateUser(ctx echo.Context, filter map[string]interface{}, request model.UpdateUsers) (result model.UpdateUsers, err error) {
	tx := repository.Postgre.Db.Model(model.UpdateUsers{}).Where(filter).Table("users").Updates(&request)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	result = request
	return
}

func (repository *usersRepositoryImpl) DeleteUser(ctx echo.Context, filter map[string]interface{}) (result model.Users, err error) {
	tx := repository.Postgre.Db.Where(filter).Delete(&result)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
