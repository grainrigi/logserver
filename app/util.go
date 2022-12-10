package app

import (
	"fmt"
	"net/http"
	"strconv"

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

func IntParam(c echo.Context, name string) (int, error) {
	i, err := strconv.Atoi(c.Param(name))

	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("path param '%s' must be a number", name))
	}

	return i, nil
}

type ErrorDTO struct {
	Error string `json:"error"`
}
