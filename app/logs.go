package app

import (
	"logserver/data"
	"logserver/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func postLog(c echo.Context) error {
	l := new(data.Log)
	if err := c.Bind(l); err != nil {
		return err
	}
	if err := db.InsertLog(c.Request().Context(), l); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getLogs(c echo.Context) error {
	logs, err := db.GetLogs(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, logs)
}
