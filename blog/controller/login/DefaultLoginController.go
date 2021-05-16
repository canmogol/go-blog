package blog

import (
	"net/http"

	serviceLogin "github.com/canmogol/godi/blog/service/login"
)

type defaultLoginController struct {
	loginService serviceLogin.LoginService
	validator    *loginRequestValidator
}

func newDefaultLoginController(loginService serviceLogin.LoginService) LoginController {
	return &defaultLoginController{
		loginService: loginService,
		validator:    &loginRequestValidator{},
	}
}

func (controller *defaultLoginController) CreateLoginRequest(username string, password string) *LoginRequest {
	return &LoginRequest{
		username: username,
		password: password,
	}
}

func (controller *defaultLoginController) Login(loginRequest *LoginRequest) *LoginResponse {
	_, err := controller.validator.validate(loginRequest)
	if err != nil {
		return controller.createLoginResponse(http.StatusBadRequest, err.Error())
	}
	logged, err := controller.loginService.Login(loginRequest.username, loginRequest.password)
	if err != nil {
		return controller.createLoginResponse(http.StatusInternalServerError, err.Error())
	}
	if !logged {
		return controller.createLoginResponse(http.StatusUnauthorized, "username and/or password are not correct.")
	}
	return controller.createLoginResponse(http.StatusOK, "logged in successfully.")
}

func (controller *defaultLoginController) createLoginResponse(httpStatusCode int, message string) *LoginResponse {
	return &LoginResponse{
		StatusCode: httpStatusCode,
		Message:    message,
	}
}
