package blog

import (
	"fmt"
	"net/http"
	"testing"
)

const internalServerErrorMessage = "Internal Server Error Message."

type successfulLoginService struct{}

func (service *successfulLoginService) Login(username string, password string) (bool, error) {
	return true, nil
}

type failureLoginService struct{}

func (service *failureLoginService) Login(username string, password string) (bool, error) {
	return false, nil
}

type errorLoginService struct{}

func (service *errorLoginService) Login(username string, password string) (bool, error) {
	return false, fmt.Errorf(internalServerErrorMessage)
}

func TestLogin_HttpLoginController(t *testing.T) {
	emptyUsernameError := "Username cannot be empty."
	emptyPasswordError := "Password cannot be empty."
	unauthorizedError := "username and/or password are not correct."
	successfulLoginMessage := "logged in successfully."
	emptyLoginService := &successfulLoginService{}
	errorLoginService := &errorLoginService{}
	failureLoginService := &failureLoginService{}
	successfulLoginService := &successfulLoginService{}
	httpLoginController := newHttpLoginController(emptyLoginService)
	httpLoginControllerError := newHttpLoginController(errorLoginService)
	httpLoginControllerFailure := newHttpLoginController(failureLoginService)
	httpLoginControllerSuccessful := newHttpLoginController(successfulLoginService)

	// Empty username
	loginRequestEmptyUsername := &LoginRequest{
		Username: "",
		Password: "NON-EMPTY-PASSWORD",
	}
	loginResponseEmptyUsername := httpLoginController.Login(loginRequestEmptyUsername)
	if loginResponseEmptyUsername.StatusCode != http.StatusBadRequest {
		t.Error("The LoginController should return a BadRequest for missing/empty username.")
	}
	if loginResponseEmptyUsername.Message != emptyUsernameError {
		t.Errorf("The LoginController should return an error message for missing username, expected: '%s', actual: '%s'.", emptyUsernameError, loginResponseEmptyUsername.Message)
	}

	// Empty password
	loginRequestEmptyPassword := &LoginRequest{
		Username: "NON-EMPTY-USERNAME",
		Password: "",
	}
	loginResponseEmptyPassword := httpLoginController.Login(loginRequestEmptyPassword)
	if loginResponseEmptyPassword.StatusCode != http.StatusBadRequest {
		t.Error("The LoginController should return a BadRequest for missing/empty password.")
	}
	if loginResponseEmptyPassword.Message != emptyPasswordError {
		t.Errorf("The LoginController should return an error message for missing password, expected: '%s', actual: '%s'.", emptyPasswordError, loginResponseEmptyPassword.Message)
	}

	// Service level Error
	loginRequestError := &LoginRequest{
		Username: "not-important",
		Password: "not-important",
	}
	loginResponseError := httpLoginControllerError.Login(loginRequestError)
	if loginResponseError.StatusCode != http.StatusInternalServerError {
		t.Error("The LoginController should return an InternalServerError when the service returns an error.")
	}
	if loginResponseError.Message != internalServerErrorMessage {
		t.Errorf("Expected InternalServerError message: '%s' is different from the actual error message: '%s'.", internalServerErrorMessage, loginResponseError.Message)
	}

	// Wrong Username and/or Password
	loginRequestFailure := &LoginRequest{
		Username: "wrong-username",
		Password: "wrong-password",
	}
	loginResponseFailure := httpLoginControllerFailure.Login(loginRequestFailure)
	if loginResponseFailure.StatusCode != http.StatusUnauthorized {
		t.Error("The LoginController should return an Unauthorized when the service returns a failure.")
	}
	if loginResponseFailure.Message != unauthorizedError {
		t.Errorf("Expected Unauthorized message: '%s' is different from the actual error message: '%s'.", unauthorizedError, loginResponseFailure.Message)
	}

	/// Correct Username and/or Password
	loginRequestSuccessful := &LoginRequest{
		Username: "correct-username",
		Password: "correct-password",
	}
	loginResponseSuccessful := httpLoginControllerSuccessful.Login(loginRequestSuccessful)
	if loginResponseSuccessful.StatusCode != http.StatusOK {
		t.Error("The LoginController should return an OK when the service returns a success.")
	}
	if loginResponseSuccessful.Message != successfulLoginMessage {
		t.Errorf("Expected OK message: '%s' is different from the actual OK message: '%s'.", successfulLoginMessage, loginResponseSuccessful.Message)
	}

	t.Log("TestLogin_HttpLoginController successful.")
}
