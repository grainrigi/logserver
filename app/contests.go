package app

import (
	"database/sql"
	"logserver/data"
	"logserver/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getContests(c echo.Context) error {
	cs, err := db.ReadContests(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cs)
}

func getContestWithLogs(c echo.Context) error {
	id, err := IntParam(c, "id")
	if err != nil {
		return err
	}

	contest, err := db.ReadContestWithLogs(c.Request().Context(), id)
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, contest)
}

func postContest(c echo.Context) error {
	contest, err := BindAndValidate[data.Contest](c)
	if err != nil {
		return err
	}

	if err := contest.Normalize(); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorDTO{Error: err.Error()})
	}

	id, err := db.InsertContest(c.Request().Context(), contest)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, struct {
		Id int `json:"id"`
	}{Id: id})
}

func putContest(c echo.Context) error {
	contest, err := BindAndValidate[data.Contest](c)
	if err != nil {
		return err
	}

	if err := contest.Normalize(); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorDTO{Error: err.Error()})
	}

	err = db.UpdateContest(c.Request().Context(), contest)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func deleteContest(c echo.Context) error {
	id, err := IntParam(c, "id")
	if err != nil {
		return err
	}

	err = db.DeleteContest(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
