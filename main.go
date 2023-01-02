package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	server := echo.New()

	server.GET("/", hello)
	server.POST("/login", login)
	server.Logger.Fatal(server.Start(":8888"))

}

type X struct {
	Text string `json:"text"`
}

func hello(c echo.Context) error {
	x := &X{
		Text: "Hello world",
	}
	return c.JSON(http.StatusOK, x)
}

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"username" xml:"password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	c.Bind(req)
	log.Printf("req data %+v", req)
	if req.Username != "admin" || req.Password != "112233" {
		return c.String(http.StatusUnauthorized, "username/password invalid")
	}

	return c.JSON(http.StatusOK, &LoginResponse{
		Token: "123456",
	})
}
