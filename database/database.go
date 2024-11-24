package database

import (
	"database/sql"
	"errors"

	l "github.com/peterszarvas94/goat/logger"
)

type Service struct {
	DB *sql.DB
}

var service = &Service{}

func StartSqliteConnection(path string) error {
	conn, err := sql.Open("sqlite", path)
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
