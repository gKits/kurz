package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gkits/kurz/internal/db"
	"github.com/gkits/kurz/internal/types"
	"github.com/gkits/kurz/internal/utils"
	"github.com/gkits/kurz/internal/views/component"
	"github.com/gkits/kurz/internal/views/page"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

func handleHome(ctx echo.Context) error {
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
	links, err := db.GetLinks(ctx.Request().Context(), q)
	if err != nil {
		ctx.Logger().Error("failed to get links: ", err)
		return err
	}
	return render(ctx, component.LinkList(links))
}

func handleCreateLink(ctx echo.Context) error {
	var link types.Link
	if err := ctx.Bind(&link); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "form binding failed")
	}
	if err := ctx.Validate(link); err != nil {
		ctx.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}

	link.CreatedAt = time.Now()
	link.CreatedBy = "joe"
	link.Ref = utils.GenerateRef()

	insertedLink, err := db.InsertLink(ctx.Request().Context(), link)
	if err != nil {
		ctx.Logger().Error("failed to insert link: ", err)
		return echo.ErrInternalServerError
	}
	return render(ctx, component.LinkListRow(insertedLink))
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
	if link.IsExpired() {
		return echo.ErrGone
	}
	return ctx.Redirect(http.StatusMovedPermanently, link.Target)
}
