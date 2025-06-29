package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Link struct {
	ID        int64     `db:"id"`
	Ref       string    `db:"ref"`
	Target    string    `db:"target"     form:"target"     validate:"url"`
	ExpiresAt NullTime  `db:"expires_at" form:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"created_by"`
}

func (l Link) IsExpired() bool { return l.ExpiresAt.Valid && l.ExpiresAt.Time.Before(time.Now()) }

type LinkQuery struct {
	Offset uint `query:"offset"`
	Limit  uint `query:"limit"`
}

type NullTime sql.NullTime

func (nt *NullTime) UnmarshalParam(param string) error {
	fmt.Println(param)
	t, err := time.Parse(time.RFC3339, param)
	nt.Time = t
	nt.Valid = err == nil
	return nil
}

func (nt *NullTime) Scan(src any) error {
	var sqln sql.NullTime
	if err := sqln.Scan(src); err != nil {
		return err
	}
	*nt = NullTime(sqln)
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	return sql.NullTime(nt).Value()
}
