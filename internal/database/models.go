// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Url struct {
	ID          uuid.UUID
	CreatedAt   sql.NullTime
	UrlID       string
	OriginalUrl string
}