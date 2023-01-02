package handler

import (
	"EchoFramework/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	x := &models.X{
		Text: "Hello world",
	}
	return c.JSON(http.StatusOK, x)
}
