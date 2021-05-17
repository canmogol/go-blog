package blog

import (
	"fmt"

	serviceLogin "github.com/canmogol/godi/blog/service/login"
)

func CreateLoginController(controllerType string, loginService serviceLogin.LoginService) (LoginController, error) {
	if controllerType == "default" {
		return newHttpLoginController(loginService), nil
	}
	return nil, fmt.Errorf("Could not create LoginController, unknown controller type passed: '%s'", controllerType)
}
