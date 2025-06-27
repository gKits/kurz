package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gkits/kurz/internal/db"
	"github.com/gkits/kurz/internal/types"
	"github.com/gkits/kurz/internal/views/component"
	"github.com/gkits/kurz/internal/views/page"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.AddTrailingSlash())

	e.GET("/", handleHome)
	e.GET("/r/:ref", handleRedirectToLink)

	e.GET("/links", handleGetLinks)
	e.POST("/links", handleCreateLink)
	e.GET("/links/:id", handleGetLink)
	e.DELETE("/links/:id", handleDeleteLink)

	return e
}

func handleHome(ctx echo.Context) error {
	fmt.Println("hello")
	return render(ctx, page.Home())
}

func handleGetLink(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	link, err := db.GetLinkByID(ctx.Request().Context(), int64(id))
	if err != nil {
		return echo.ErrInternalServerError
	}
	return render(ctx, component.LinkListRow(link))
}

func handleGetLinks(ctx echo.Context) error {
	var q types.LinkQuery
	if err := schema.NewDecoder().Decode(&q, ctx.QueryParams()); err != nil {
		return echo.ErrBadRequest
	}
	links, err := db.GetLinks(ctx.Request().Context(), types.LinkQuery{})
	if err != nil {
		return err
	}
	return render(ctx, component.LinkList(links))
}

func handleCreateLink(ctx echo.Context) error {
	var linkCreate types.LinkCreate
	if err := ctx.Bind(&linkCreate); err != nil {
		fmt.Println(err)
		return echo.ErrBadRequest
	}
	fmt.Println(linkCreate)
	return nil
}

func handleDeleteLink(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := db.DeleteLink(ctx.Request().Context(), id); err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.NoContent(http.StatusNoContent)
}

func handleRedirectToLink(ctx echo.Context) error {
	ref := ctx.Param("ref")
	link, err := db.GetLinkByRef(ctx.Request().Context(), ref)
	if err != nil {
		return err
	}
	return ctx.Redirect(http.StatusMovedPermanently, link.Target)
}

func render(ctx echo.Context, component templ.Component) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}
