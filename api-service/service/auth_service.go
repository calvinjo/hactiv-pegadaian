package service

import (
	"api-service/config"
	"api-service/model"
	"api-service/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	ProcessRegister(ctx echo.Context, request model.RequestRegister) (message string, statusCode string, data model.ResponseRegister)
	ProcessLogin(ctx echo.Context, request model.RequestLogin) (message string, statusCode string, data model.ResponseLogin)
}

type authServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewAuthService(newUsersRepository repository.UsersRepository) AuthService {
	return &authServiceImpl{
		UsersRepository: newUsersRepository,
	}
}

func (service *authServiceImpl) ProcessRegister(ctx echo.Context, request model.RequestRegister) (message string, statusCode string, data model.ResponseRegister) {
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

	resultDetailUser := service.UsersRepository.DetailUser(ctx, model.RepoRequestDetailUser{
		Filter: map[string]interface{}{
			"username": reqCreateUser.Username,
		},
	})

	if resultDetailUser.IsError {
		message = resultDetailUser.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	if resultDetailUser.Data.UserID > 0 {
		message = "The User Has Been Registered"
		statusCode = config.Failed

		return
	}

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

func (service *authServiceImpl) ProcessLogin(ctx echo.Context, request model.RequestLogin) (message string, statusCode string, data model.ResponseLogin) {
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

	detailUser := service.UsersRepository.DetailUser(ctx, model.RepoRequestDetailUser{
		Filter: map[string]interface{}{
			"username": request.Username,
		},
	})

	if detailUser.IsNotFound {
		message = "Data Not Found"
		statusCode = config.DataNotFound
	}

	if detailUser.IsError {
		message = detailUser.ErrorMessage.Error()
		statusCode = config.Failed

		return
	}

	// Check password
	errPass := bcrypt.CompareHashAndPassword([]byte(detailUser.Data.Password), []byte(request.Password))
	if errPass != nil {
		message = "Password Invalid"
		statusCode = config.Failed

		return
	}

	claims := &model.JwtClaims{
		detailUser.Data.UserID,
		detailUser.Data.Roles,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("mysecret"))

	data.Token = signedToken
	data.UserID = detailUser.Data.UserID
	data.Roles = detailUser.Data.Roles
	return
}
