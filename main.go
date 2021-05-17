package main

import (
	blog "github.com/canmogol/godi/blog"
	login "github.com/canmogol/godi/blog/controller/login"
	"github.com/labstack/echo/v4"
)

func main() {
	blogApp := blog.New()
	e := echo.New()

	e.POST("/login", func(c echo.Context) error {
		u := c.FormValue("username")
		p := c.FormValue("password")
		loginRequest := &login.LoginRequest{
			Username: u,
			Password: p,
		}
		loginResponse := blogApp.LoginController.Login(loginRequest)
		return c.String(loginResponse.StatusCode, loginResponse.Message)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
