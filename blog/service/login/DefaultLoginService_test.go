package blog

import (
	"testing"
)

type emptyLoginRepository struct{}

func (empty *emptyLoginRepository) Login(username string, password string) bool {
	return true
}

func TestFindAll_DefaultLoginService(t *testing.T) {
	emptyString := ""
	username := "username"
	password := "password"
	loginRepository := &emptyLoginRepository{}

	defaultLoginService := newDefaultLoginService(loginRepository)
	if defaultLoginService == nil {
		t.Error("DefaultLoginService cannot be nil.")
	}

	_, errorEmpty := defaultLoginService.Login(emptyString, emptyString)
	if errorEmpty == nil {
		t.Error("LoginService should not allow login with empty string for either username or password.")
	}

	_, errorEmptyPassword := defaultLoginService.Login(username, emptyString)
	if errorEmptyPassword == nil {
		t.Error("LoginService should not allow login with empty password.")
	}

	_, errorEmptyUsername := defaultLoginService.Login(emptyString, password)
	if errorEmptyUsername == nil {
		t.Error("LoginService should not allow login with empty username.")
	}

	logged, errorNil := defaultLoginService.Login(username, password)
	if errorNil != nil {
		t.Errorf("LoginService should allow logins with username and password, error: '%s'.", errorNil.Error())
	}
	if !logged {
		t.Error("LoginService should allow logins with username and password, Login method returned false.")
	}

	t.Log("TestFindAll_DefaultLoginService successful.")
}
