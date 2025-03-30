package helpers

import (
	"fmt"
	"strings"
)

func generateMigrationSql(modelname string, create bool) string {
	if !create {
		return `-- +goose Up

-- +goose Down`
	}

	return fmt.Sprintf(`-- +goose Up
CREATE TABLE %s (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE %s;`, modelname, modelname)
}

func generateSchemaSql(modelname string) string {
	return fmt.Sprintf(`CREATE TABLE %s (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`,
		modelname,
	)
}

func generateQueriesSql(modelname string) string {
	uppercasemodelName := fmt.Sprintf("%s%s", strings.ToUpper(string(modelname[0])), modelname[1:])
	return fmt.Sprintf(`-- name: Get%sByID :one
SELECT *
FROM %s
WHERE id = ?;

-- name: List%ss :many
SELECT *
FROM %s
ORDER BY name;

-- name: Create%s :one
INSERT INTO %s (id, name)
VALUES (?, ?)
RETURNING *;

-- name: Update%s :one
UPDATE %s
SET name = ?
WHERE id = ?
RETURNING *;

-- name: Delete%s :exec
DELETE FROM %s
WHERE id = ?;`,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
	)
}
