package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Normalizable interface {
	Normalize() error
}

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

func BindAndNormalize[T any, PT interface {
	Normalizable // ポインタ型がNormalizableであることを宣言(値型は(T)Normalize()を実装しないからNormalizableではない=any)
	*T
}](c echo.Context) (*T, error) {
	var i PT
	i, err := BindAndValidate[T](c)
	if err != nil {
		return nil, err
	}
	if err := i.Normalize(); err != nil {
		return nil, c.JSON(http.StatusBadRequest, ErrorDTO{Error: err.Error()})
	}
	return (*T)(i), nil
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

type IDDTO struct {
	ID int `json:"id"`
}
