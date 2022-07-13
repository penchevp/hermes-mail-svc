package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const dataSourceNamePlaceholder = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

var dbConn *gorm.DB

type Config struct {
	Host     string
	Port     int32
	User     string
	Password string
	Catalog  string
}

func InitialiseConnection(config Config) error {
	if dbConn != nil {
		// close off existing dbConn
		sqlDB, err := dbConn.DB()
		if err != nil {
			return err
		}

		if err := sqlDB.Close(); err != nil {
			return err
		}
	}

	dataSourceName := fmt.Sprintf(dataSourceNamePlaceholder, config.Host, config.User, config.Password, config.Catalog, config.Port)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Customer{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	dbConn = db

	return nil
}

func GetConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}

	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}

	return dbConn, nil
}

func CloseConnection() error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
