package blog

import (
	"testing"
)

func TestLogin_DefaultLoginRepository(t *testing.T) {
	loginRepositoryDefault := newDefaultLoginRepository()
	loggedEmpty := loginRepositoryDefault.Login("", "")
	if loggedEmpty == false {
		t.Error("The default LoginRepository should always return true.")
	}
	t.Log("TestLogin successful.")
}
