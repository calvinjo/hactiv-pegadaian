package config

import (
	"fmt"
	"transaction-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSql struct {
	Db *gorm.DB
}

func NewPostgreDatabase(env model.Args) *PostgreSql {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.PostgreHost, env.PostgreUser, env.PostgrePassword, env.PostgreDb, env.PostgrePort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}

	//Migrate
	db.AutoMigrate(&model.Loan{}, &model.Installment{}, &model.MaxLoan{})

	return &PostgreSql{Db: db}
}
