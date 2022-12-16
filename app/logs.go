package app

import (
	"logserver/data"
	"logserver/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func postLog(c echo.Context) error {
	l, err := BindAndNormalize[data.Log](c)
	if err != nil {
		return err
	}

	id, err := db.InsertLog(c.Request().Context(), l)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, IDDTO{ID: id})
}

func getLogs(c echo.Context) error {
	cid, err := IntParam(c, "cid")
	if err != nil {
		return err
	}

	logs, err := db.ReadLogs(c.Request().Context(), cid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, logs)
}

func putLog(c echo.Context) error {
	l, err := BindAndNormalize[data.Log](c)
	if err != nil {
		return err
	}

	if err := db.UpdateLog(c.Request().Context(), l); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func deleteLog(c echo.Context) error {
	id, err := IntParam(c, "id")
	if err != nil {
		return err
	}

	if err := db.DeleteLog(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
