package database

import (
	"errors"

	l "github.com/peterszarvas94/goat/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

var service = &Service{}

func StartSqliteConnection(path string) error {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		l.Logger.Error(err.Error())
		return err
	}

	service.DB = conn

	l.Logger.Debug("DB setup is done")
	return nil
}

func Get() (*Service, error) {
	if service == nil {
		err := errors.New("Database is not yet set up")
		l.Logger.Error(err.Error())
		return nil, err
	}
	return service, nil
}
