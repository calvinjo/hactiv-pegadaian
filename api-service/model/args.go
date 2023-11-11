package model

type Args struct {
	AppsPort               string `mapstructure:"APPS_PORT" validate:"required"`
	TransactionsServiceUrl string `mapstructure:"TRANSACTION_SERVICE_URL" validate:"required"`
	UsersServiceUrl        string `mapstructure:"USERS_SERVICE_URL" validate:"required"`
}
