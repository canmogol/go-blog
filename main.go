package main

import (
	blog "github.com/canmogol/godi/blog"
	"github.com/labstack/echo/v4"
)

func main() {
	blogApp := blog.New()
	e := echo.New()

	e.POST("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		loginRequest := blogApp.LoginController.CreateLoginRequest(username, password)
		loginResponse := blogApp.LoginController.Login(loginRequest)
		return c.String(loginResponse.StatusCode, loginResponse.Message)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
