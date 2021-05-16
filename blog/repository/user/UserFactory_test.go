package blog

import (
	"testing"
)

func TestCreateUserRepository(t *testing.T) {
	const expectedErrorMessage = "Could not create UserRepository, unknown repository type passed: ''"
	const emptyString = ""
	const defaultRepositoryType = "default"
	expectedUserRepository := newUserRepository()
	userRepositoryEmpty, err := CreateUserRepository(emptyString)
	if userRepositoryEmpty != nil {
		t.Error("There cannot be a UserRepository with empty repositoryType.")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("The expected error message: '%s' is different from the actual message: '%s'", expectedErrorMessage, err.Error())
	}

	userRepositoryDefault, err := CreateUserRepository(defaultRepositoryType)
	if userRepositoryDefault == nil {
		t.Errorf("There has to be a 'default' UserRepository, provided repositoryType: '%s'.", defaultRepositoryType)
	}
	if err != nil {
		t.Errorf("There cannot be an error for default repositoryType: '%s', error: '%s'.", defaultRepositoryType, err.Error())
	}
	if userRepositoryDefault != expectedUserRepository {
		t.Error("The default UserRepository is not the expected UserRepository.")
	}
	t.Logf("TestCreateUserRepository successful %s.", err)
}
