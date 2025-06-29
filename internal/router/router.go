package router

import (
	"os"

	"github.com/a-h/templ"
	fs "github.com/gkits/kurz"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ziflex/lecho/v3"

	"github.com/go-playground/validator"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	log := lecho.New(
		os.Stdout,
		lecho.WithLevel(log.DEBUG),
		lecho.WithTimestamp(),
	)
	e.Logger = log

	e.Validator = &eValidator{validator.New()}

	e.Use(middleware.RequestID())
	e.Use(lecho.Middleware(lecho.Config{
		Logger: log,
	}))
	e.Use(middleware.AddTrailingSlash())

	e.GET("/", handleHome)
	e.GET("/r/:ref", handleRedirectToLink)
	e.StaticFS("/public", echo.MustSubFS(fs.Public(), "public"))

	e.GET("/links", handleGetLinks)
	e.POST("/links", handleCreateLink)
	e.GET("/links/:id", handleGetLink)
	e.DELETE("/links/:id", handleDeleteLink)

	return e
}

func render(ctx echo.Context, component templ.Component) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}
