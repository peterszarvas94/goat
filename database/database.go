package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type connectionT struct {
	DB *gorm.DB
}

var connection = &connectionT{}

func StartSqliteConnection(path string) error {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return err
	}

	connection.DB = conn

	return nil
}

func Get() (*connectionT, error) {
	if connection == nil {
		return nil, errors.New("Database is not yet set up")
	}
	return connection, nil
}
