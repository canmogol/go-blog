package blog

import (
	"testing"
)

type loginRepositoryEmpty struct{}

func (loginRepositoryEmpty *loginRepositoryEmpty) Login(username string, password string) bool {
	return true
}

func TestCreateLoginRepository(t *testing.T) {
	expectedErrorMessage := "Could not create LoginService, unknown service type passed: ''"
	emptyString := ""
	loginRepositoryEmpty := &loginRepositoryEmpty{}

	loginServiceEmpty, err := CreateLoginService(emptyString, loginRepositoryEmpty)
	if loginServiceEmpty != nil {
		t.Error("There cannot be a LoginService with empty repositoryType.")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("The expected error message: '%s' is different from the actual message: '%s'", expectedErrorMessage, err.Error())
	}

	defaultServiceType := "default"
	loginServiceDefault, err := CreateLoginService(defaultServiceType, loginRepositoryEmpty)
	if loginServiceDefault == nil {
		t.Errorf("There has to be a 'default' LoginService, provided serviceType: '%s'.", defaultServiceType)
	}
	if err != nil {
		t.Errorf("There cannot be an error for default serviceType: '%s', error: '%s'.", defaultServiceType, err.Error())
	}
	if loginServiceDefault == nil {
		t.Error("The default LoginService cannot be nil.")
	}
	t.Logf("TestCreateLoginRepository successful %s.", err)
}
