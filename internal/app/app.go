package app

import (
	"github.com/gkits/kurz/internal/router"
	"github.com/labstack/echo/v4"
)

type App struct {
	e *echo.Echo
}

func New() *App {
	app := &App{
		e: router.New(),
	}
	return app
}

func (a *App) Run() error {
	return a.e.Start(":4000")
}
