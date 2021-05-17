package blog

import (
	"fmt"

	repositoryLogin "github.com/canmogol/godi/blog/repository/login"
)

func CreateLoginService(serviceType string, loginRepository repositoryLogin.LoginRepository) (LoginService, error) {
	if serviceType == "default" {
		return newDefaultLoginService(loginRepository), nil
	}
	return nil, fmt.Errorf("Could not create LoginService, unknown service type passed: '%s'", serviceType)
}
