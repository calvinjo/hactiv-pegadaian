package config

import (
	"log"
	"users-service/model"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Env interface {
	SetArgs() model.Args
}

type envImpl struct{}

func (env *envImpl) SetArgs() model.Args {
	var args model.Args

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}

	if err := viper.Unmarshal(&args); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	validate := validator.New()
	if err := validate.Struct(args); err != nil {
		log.Fatalf("Validation error: %s", err)
	}

	return args
}

func New() Env {
	return &envImpl{}
}
