package model

type Args struct {
	AppsPort        string `mapstructure:"APPS_PORT" validate:"required"`
	PostgreHost     string `mapstructure:"POSTGRE_HOST" validate:"required"`
	PostgreUser     string `mapstructure:"POSTGRE_USER" validate:"required"`
	PostgrePassword string `mapstructure:"POSTGRE_PASSWORD" validate:"required"`
	PostgreDb       string `mapstructure:"POSTGRE_DB" validate:"required"`
	PostgrePort     string `mapstructure:"POSTGRE_PORT" validate:"required"`
}
