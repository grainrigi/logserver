package app

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var validate = validator.New()

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/contests", getContests)
	e.GET("/contests/:id", getContest)
	e.POST("/contests", postContest)
	e.PUT("/contests/:id", putContest)
	e.DELETE("/contests/:id", deleteContest)

	e.GET("/contests/:cid/logs", getLogs)
	e.POST("/contests/:cid/logs", postLog)
	e.PUT("/contests/:cid/logs/:id", putLog)
	e.DELETE("/contests/:cid/logs/:id", deleteLog)

	e.GET("/operators", getOperators)
	e.POST("/operators", postOperator)
	e.DELETE("/operators/:id", deleteOperator)
	e.PUT("/operators/:id", putOperator)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "3030"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
