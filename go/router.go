package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func api(c echo.Context) error {

	return c.String(http.StatusOK, "test")
}
