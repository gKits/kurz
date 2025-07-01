package router

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type eValidator struct {
	valid *validator.Validate
}

func (valid *eValidator) Validate(i any) error {
	if err := valid.valid.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
