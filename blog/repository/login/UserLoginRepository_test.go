package blog

import (
	"fmt"
	"testing"

	repositoryUser "github.com/canmogol/godi/blog/repository/user"
)

type userRepositorySuccess struct{}

func (mock *userRepositorySuccess) FindUserWithUsernameAndPassword(username string, password string) (*repositoryUser.UserEntity, error) {
	return &repositoryUser.UserEntity{}, nil
}

func (mock *userRepositorySuccess) FindAll() []*repositoryUser.UserEntity {
	return []*repositoryUser.UserEntity{}
}

type userRepositoryError struct{}

func (mock *userRepositoryError) FindUserWithUsernameAndPassword(username string, password string) (*repositoryUser.UserEntity, error) {
	return nil, fmt.Errorf("No User Found")
}

func (mock *userRepositoryError) FindAll() []*repositoryUser.UserEntity {
	return []*repositoryUser.UserEntity{}
}

func TestLogin_ForUserLoginRepository(t *testing.T) {
	userRepositoryError := &userRepositoryError{}
	loginRepositoryUserError := newUserLoginRepository(userRepositoryError)
	loggedFalse := loginRepositoryUserError.Login("", "")
	if loggedFalse != false {
		t.Error("The UserRepository did not return a user, login method should fail.")
	}

	userRepositorySuccess := &userRepositorySuccess{}
	loginRepositoryUser := newUserLoginRepository(userRepositorySuccess)
	loggedTrue := loginRepositoryUser.Login("", "")
	if loggedTrue != true {
		t.Error("The UserRepository returned a valid user, login method should be successful.")
	}
	t.Log("TestLogin successful.")
}
