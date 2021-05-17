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

func (controller *httpLoginController) CreateLoginRequest(username string, password string) *LoginRequest {
	return &LoginRequest{
		username: username,
		password: password,
	}
}

func (controller *httpLoginController) Login(loginRequest *LoginRequest) *LoginResponse {
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

func (controller *httpLoginController) createLoginResponse(httpStatusCode int, message string) *LoginResponse {
	return &LoginResponse{
		StatusCode: httpStatusCode,
		Message:    message,
	}
}
