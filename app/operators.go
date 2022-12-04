package app

import (
	"logserver/data"
	"logserver/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getOperators(c echo.Context) error {
	ops, err := db.ReadOperators(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ops)
}

func postOperator(c echo.Context) error {
	op, err := BindAndValidate[data.Operator](c)
	if err != nil {
		return err
	}

	if err := db.InsertOperator(c.Request().Context(), op); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func deleteOperator(c echo.Context) error {
	var op data.Operator

	if err := c.Bind(&op); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path")
	}

	if err := db.DeleteOperator(c.Request().Context(), op.ID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func putOperator(c echo.Context) error {
	op, err := BindAndValidate[data.Operator](c)
	if err != nil {
		return err
	}

	if err := db.UpdateOperator(c.Request().Context(), op); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
