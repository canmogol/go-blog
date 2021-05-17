package blog

import (
	"fmt"
	"strings"
)

type loginRequestValidator struct {
}

func (validator *loginRequestValidator) validate(loginRequest *LoginRequest) (bool, error) {
	if strings.TrimSpace(loginRequest.Username) == "" {
		return false, fmt.Errorf("Username cannot be empty.")
	}
	if strings.TrimSpace(loginRequest.Password) == "" {
		return false, fmt.Errorf("Password cannot be empty.")
	}
	return true, nil
}
