package main

import (
	"EchoFramework/handler"
	"EchoFramework/midware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())

	server.GET("/", handler.Hello)
	server.POST("/login", handler.Login, middleware.BasicAuth(midware.BasicAuth))
	server.Logger.Fatal(server.Start(":8888"))

}
