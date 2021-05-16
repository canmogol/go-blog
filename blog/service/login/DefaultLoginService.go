package blog

import (
	"fmt"

	repositoryLogin "github.com/canmogol/godi/blog/repository/login"
)

type defaultLoginService struct {
	loginRepository repositoryLogin.LoginRepository
}

func (service *defaultLoginService) Login(username string, password string) (bool, error) {
	if username == "" && password == "" {
		return false, fmt.Errorf("username and/or password cannot be empty.")
	}
	return service.loginRepository.Login(username, password), nil
}

func newDefaultLoginService(loginRepository repositoryLogin.LoginRepository) LoginService {
	return &defaultLoginService{
		loginRepository: loginRepository,
	}
}
