package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/peterszarvas94/goat/pkg/logger"
)

var db *sql.DB

func Connect(path string) (*sql.DB, error) {
	logger.Debug("Connecting to %s", path)
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	db = conn

	logger.Debug("DB setup is done")
	return db, nil
}

func Get() (*sql.DB, error) {
	if db == nil {
		err := errors.New("Database is not yet set up")
		logger.Error(err.Error())
		return nil, err
	}
	return db, nil
}
