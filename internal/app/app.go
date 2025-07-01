package app

import (
	"fmt"

	"github.com/gkits/kurz/config"
	"github.com/gkits/kurz/internal/db"
	"github.com/gkits/kurz/internal/router"
	"github.com/labstack/echo/v4"
)

type App struct {
	e *echo.Echo
}

func New(cfg config.Config) (*App, error) {
	conn, err := cfg.Database.Open()
	if err != nil {
		return nil, fmt.Errorf("app: failed to open db connection: %w", err)
	}
	db.Init(conn, cfg.Database.Driver.String())

	app := &App{
		e: router.New(),
	}
	return app, nil
}

func (a *App) Run() error {
	return a.e.Start(":4000")
}
