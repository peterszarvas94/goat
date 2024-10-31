package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func OpenTurso(url, token string) (*sql.DB, error) {
	dbUrl := fmt.Sprintf("%s?authToken=%s", url, token)
	conn, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, err
	}

	DB = conn

	return conn, nil
}

func Connect() (*sql.DB, error) {
	if DB == nil {
		return nil, errors.New("Database is not yet set up")
	}
	return DB, nil
}
