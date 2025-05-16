package database

import (
	"database/sql"
	"errors"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect(path string) (*sql.DB, error) {
	slog.Debug("Connecting", "path", path)
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	db = conn

	slog.Debug("DB setup is done")
	return db, nil
}

func Get() (*sql.DB, error) {
	if db == nil {
		err := errors.New("Database is not yet set up")
		slog.Error(err.Error())
		return nil, err
	}
	return db, nil
}
