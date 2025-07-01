package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Driver   DatabaseDriver `env:"DRIVER"   default:"sqlite3"`
	Host     string         `env:"HOST"     default:"./kurz.db"`
	Port     uint16         `env:"PORT"`
	Username string         `env:"USERNAME"`
	Password string         `env:"PASSWORD"`
}

func (db Database) DSN() string {
	switch db.Driver {
	case SQLite3:
		return db.Host
	default:
	}
	return ""
}

func (db Database) Open() (*sql.DB, error) {
	return sql.Open(db.Driver.String(), db.DSN())
}

type DatabaseDriver string

const (
	SQLite3 = "sqlite3"
)

func (dbd DatabaseDriver) String() string {
	return string(dbd)
}

func (dbd *DatabaseDriver) UnmarshalText(text []byte) error {
	switch DatabaseDriver(text) {
	case SQLite3:
		*dbd = DatabaseDriver(text)
	default:
		return fmt.Errorf("config: unknow database driver '%s': has to be one of (sqlite)", text)
	}
	return nil
}
