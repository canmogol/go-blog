package blog

import (
	"strings"
	"testing"
)

func TestFindAll_DefaultUserRepository(t *testing.T) {
	userRepository := newUserRepository()
	userEntityList := userRepository.FindAll()
	if userEntityList == nil {
		t.Error("UserEntity list cannot be empty.")
	}
	if len(userEntityList) == 0 {
		t.Error("UserEntity list is empty, default implementation should have at least one UserEntity.")
	}
	userEntity := userEntityList[0]
	if strings.TrimSpace(userEntity.username) == "" {
		t.Error("Username cannot be empty.")
	}
	if strings.TrimSpace(userEntity.password) == "" {
		t.Error("Password cannot be empty.")
	}
	t.Log("TestFindAll successful.")
}

func TestFindUserWithUsernameAndPassword_DefaultUserRepository(t *testing.T) {
	const expectedErrorMessage = "No user found with this username and password."
	const expectedUsername = "user123"
	const expectedPassword = "pass123"
	var nilString string
	userRepository := newUserRepository()
	userEntityNil, err := userRepository.FindUserWithUsernameAndPassword(nilString, nilString)
	if userEntityNil != nil {
		t.Error("There cannot be a UserEntity with nil username and/or password.")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("The expected the error message: '%s' and actual message: '%s' are different.", expectedErrorMessage, err.Error())
	}

	userEntityEmpty, err := userRepository.FindUserWithUsernameAndPassword("", "")
	if userEntityEmpty != nil {
		t.Error("There cannot be a UserEntity with an empty username and password.")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("The expected the error message: '%s' and actual message: '%s' are different.", expectedErrorMessage, err.Error())
	}
	userEntity, err := userRepository.FindUserWithUsernameAndPassword(expectedUsername, expectedPassword)
	if err != nil {
		t.Errorf("Default UserRepository should have a user with username: '%s' and password: '%s', instead we got error: '%s'", expectedUsername, expectedPassword, err.Error())
	}
	if userEntity.username != expectedUsername {
		t.Errorf("Expected username: '%s' is different then the actual username: '%s'", expectedUsername, userEntity.username)
	}
	if userEntity.password != expectedPassword {
		t.Errorf("Expected password: '%s' is different then the actual password: '%s'", expectedPassword, userEntity.password)
	}
	t.Log("TestFindUserWithUsernameAndPassword successful.")
}
