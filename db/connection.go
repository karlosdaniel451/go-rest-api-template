package db

import (
	"fmt"
	"github.com/karlosdaniel451/go-rest-api-template/config"
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Connect(appConfig config.AppConfig) error {
	var err error

	connectionConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		appConfig.DatabaseHost, appConfig.DatabaseUser, appConfig.DatabasePassword,
		appConfig.DatabaseName, appConfig.DatabasePort,
	)

	db, err = gorm.Open(postgres.Open(connectionConfig), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("error when initializing session to database: %s", err)
	}

	err = db.AutoMigrate(&model.Task{}, &model.User{})
	if err != nil {
		return fmt.Errorf("error when running database migrations: %s", err)
	}

	return nil
}
