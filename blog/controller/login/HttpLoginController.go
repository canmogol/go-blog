package blog

import (
	"net/http"

	serviceLogin "github.com/canmogol/godi/blog/service/login"
)

type httpLoginController struct {
	loginService serviceLogin.LoginService
	validator    *loginRequestValidator
}

func newHttpLoginController(loginService serviceLogin.LoginService) LoginController {
	return &httpLoginController{
		loginService: loginService,
		validator:    &loginRequestValidator{},
	}
}

func (controller *httpLoginController) Login(loginRequest *LoginRequest) *LoginResponse {
	_, err := controller.validator.validate(loginRequest)
	if err != nil {
		return controller.createLoginResponse(http.StatusBadRequest, err.Error())
	}
	logged, err := controller.loginService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return controller.createLoginResponse(http.StatusInternalServerError, err.Error())
	}
	if !logged {
		return controller.createLoginResponse(http.StatusUnauthorized, "username and/or password are not correct.")
	}
	return controller.createLoginResponse(http.StatusOK, "logged in successfully.")
}

func (controller *httpLoginController) createLoginResponse(httpStatusCode int, message string) *LoginResponse {
	return &LoginResponse{
		StatusCode: httpStatusCode,
		Message:    message,
	}
}
