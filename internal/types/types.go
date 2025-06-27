package types

import (
	"database/sql"
	"time"
)

type Link struct {
	ID        int64
	Ref       string
	Target    string
	ExpiresAt sql.NullTime
	CreatedAt time.Time
	CreatedBy string
}

type LinkCreate struct {
	Target    string       `form:"target"`
	ExpiresAt sql.NullTime `form:"expires_at"`
	CreatedBy string
}

type LinkQuery struct {
	Offset uint `query:"offset"`
	Limit  uint `query:"limit"`
}

type NullTime sql.NullTime
