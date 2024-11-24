package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	l "github.com/peterszarvas94/goat/logger"
)

var db *sql.DB

func StartSqliteConnection(path string) error {
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		l.Logger.Error(err.Error())
		return err
	}

	db = conn

	l.Logger.Debug("DB setup is done")
	return nil
}

func Get() (*sql.DB, error) {
	if db == nil {
		err := errors.New("Database is not yet set up")
		l.Logger.Error(err.Error())
		return nil, err
	}
	return db, nil
}
