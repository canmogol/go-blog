package blog

import (
	"fmt"

	controllerLogin "github.com/canmogol/godi/blog/controller/login"
	repositoryLogin "github.com/canmogol/godi/blog/repository/login"
	repositoryUser "github.com/canmogol/godi/blog/repository/user"
	serviceLogin "github.com/canmogol/godi/blog/service/login"
)

type BlogApplication struct {
	LoginController controllerLogin.LoginController
}

func New() *BlogApplication {
	blogApplication := &BlogApplication{}
	blogApplication.build()
	return blogApplication
}

func (app *BlogApplication) build() {
	fmt.Println("Blog application started.")
	userRepository, err := repositoryUser.CreateUserRepository("default")
	if err != nil {
		panic(err)
	}
	loginRepository, err := repositoryLogin.CreateLoginRepository("user", userRepository)
	if err != nil {
		panic(err)
	}
	loginService, err := serviceLogin.CreateLoginService("default", loginRepository)
	if err != nil {
		panic(err)
	}
	loginController, err := controllerLogin.CreateLoginController("default", loginService)
	if err != nil {
		panic(err)
	}
	app.LoginController = loginController
}
