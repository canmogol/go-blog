package blog

import (
	"testing"
)

type loginServiceEmpty struct{}

func (service *loginServiceEmpty) Login(username string, password string) (bool, error) {
	return true, nil
}

func TestCreateLoginController(t *testing.T) {
	expectedErrorMessage := "Could not create LoginController, unknown controller type passed: ''"
	emptyString := ""
	loginServiceEmpty := &loginServiceEmpty{}

	loginController, err := CreateLoginController(emptyString, loginServiceEmpty)
	if loginController != nil {
		t.Error("There cannot be a LoginController with empty controllerType.")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("The expected error message: '%s' is different from the actual message: '%s'", expectedErrorMessage, err.Error())
	}

	defaultControllerType := "default"
	loginControllerDefault, err := CreateLoginController(defaultControllerType, loginServiceEmpty)
	if loginControllerDefault == nil {
		t.Errorf("There has to be a 'default' LoginController, provided controllerType: '%s'.", defaultControllerType)
	}
	if err != nil {
		t.Errorf("There cannot be an error for default controllerType: '%s', error: '%s'.", defaultControllerType, err.Error())
	}
	if loginControllerDefault == nil {
		t.Error("The default LoginController cannot be nil.")
	}
	t.Logf("TestCreateLoginController successful %s.", err)
}
