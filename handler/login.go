package handler

import (
	"EchoFramework/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

/*func login(c echo.Context) error {
		req := new(LoginRequest)
		c.Bind(req)
		log.Printf("req data %+v", req)
		if req.Username != "admin" || req.Password != "112233" {
			return c.String(http.StatusUnauthorized, "username/password invalid")
		}


	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: "123456",
	})
}*/

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: "123456",
	})
}
