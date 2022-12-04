package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BindAndValidate[T any](c echo.Context) (*T, error) {
	i := new(T)
	if err := c.Bind(i); err != nil {
		c.Logger().Errorf("Failed to bind body: %s", err)
		return i, c.String(http.StatusBadRequest, "bad request")
	}
	if err := validate.Struct(i); err != nil {
		return i, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return i, nil
}
