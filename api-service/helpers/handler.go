package helpers

import (
	"api-service/config"
	"api-service/model"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GenerateResponse(ctx echo.Context, responseCode, additionalError string, data interface{}, err error) error {
	var message string
	if additionalError == "" {
		message = config.MessageResponse(responseCode) + additionalError
	} else {
		message = additionalError
	}

	if responseCode != config.Success {
		data = nil
	}

	response := model.Response{
		ResponseCode: responseCode,
		Message:      message,
		Data:         data,
	}

	if err != nil {
		log.Println(err)
	}

	httpCode, err := strconv.Atoi(responseCode)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return ctx.JSON(httpCode, response)
}
